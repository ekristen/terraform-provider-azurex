// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	subscriptionSettings "github.com/ekristen/terraform-provider-azurex/internal/azure/subscriptions"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SubscriptionTagsResource{}
var _ resource.ResourceWithImportState = &SubscriptionTagsResource{}

func NewSubscriptionTagsResource() resource.Resource {
	return &SubscriptionTagsResource{}
}

// SubscriptionTagsResource defines the resource implementation.
type SubscriptionTagsResource struct {
	SettingsClient      *subscriptionSettings.SettingsClient
	SubscriptionsClient *armsubscriptions.SubscriptionClient
	TagsClient          *armresources.TagsClient
	SubscriptionID      string
}

// SubscriptionTagsResourceModel describes the resource data model.
type SubscriptionTagsResourceModel struct {
	Tags              types.Map  `tfsdk:"tags"`
	InheritTags       types.Bool `tfsdk:"inherit_tags"`
	PreferContainers  types.Bool `tfsdk:"prefer_containers"`
	RemoveTags        types.Bool `tfsdk:"ondelete_remove_tags"`
	RemoteInheritTags types.Bool `tfsdk:"ondelete_remove_inherit_tags"`
}

func (r *SubscriptionTagsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscription_tags"
}

func (r *SubscriptionTagsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Subscription Tags",

		Attributes: map[string]schema.Attribute{
			"tags": schema.MapAttribute{
				Required:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Tags to apply to a subscription",
			},
			"inherit_tags": schema.BoolAttribute{
				MarkdownDescription: "Enables Inherit Tags (does not disable inherit tags on destroy)",
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"prefer_containers": schema.BoolAttribute{
				MarkdownDescription: "Prefer subscription/resource group tags over resource tags when there's a conflict",
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ondelete_remove_tags": schema.BoolAttribute{
				MarkdownDescription: "Remove tags on delete of resource",
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ondelete_remove_inherit_tags": schema.BoolAttribute{
				MarkdownDescription: "Remove tag inheritance on resource deletion",
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *SubscriptionTagsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(AzurexContext)
	if !ok {
		resp.Diagnostics.AddError("unable to obtain provider data", "provider data not available")
		return
	}

	settingsClient, err := subscriptionSettings.NewSettingsClient(data.SubscriptionID, data.IdentityCreds, nil)
	if err != nil {
		resp.Diagnostics.AddError("unable to configure settings client", fmt.Sprintf("got: %s", err.Error()))
		return
	}
	r.SettingsClient = settingsClient

	subClient, err := armsubscriptions.NewSubscriptionClient(data.IdentityCreds, nil)
	if err != nil {
		resp.Diagnostics.AddError("unable to configure subscription client", fmt.Sprintf("got: %s", err.Error()))
		return
	}
	r.SubscriptionsClient = subClient

	tagsClient, err := armresources.NewTagsClient(data.SubscriptionID, data.IdentityCreds, nil)
	if err != nil {
		resp.Diagnostics.AddError("unable to configure tags client", fmt.Sprintf("got: %s", err.Error()))
		return
	}
	r.TagsClient = tagsClient

	r.SubscriptionID = data.SubscriptionID
}

func (r *SubscriptionTagsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SubscriptionTagsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "creating subscription tags resource")

	tfTags := make(map[string]string)
	resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tfTags, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.applyTags(ctx, r.SubscriptionID, tfTags)
	if err != nil {
		resp.Diagnostics.AddError("Error applying tags to subscription", err.Error())
		return
	}

	if data.InheritTags.ValueBool() {
		tagInheritance, err := r.SettingsClient.EnableTagInheritance(ctx, data.PreferContainers.ValueBool())
		if err != nil {
			resp.Diagnostics.AddError("Error configuring tag inheritance", err.Error())
			return
		}

		if tagInheritance.Id != "" {
			data.InheritTags = types.BoolValue(true)
			data.PreferContainers = types.BoolValue(tagInheritance.Properties.PreferContainerTags)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriptionTagsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SubscriptionTagsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	scope := fmt.Sprintf("/subscriptions/%s", r.SubscriptionID)

	// Get tags using TagsClient instead of SubscriptionClient
	tagsResponse, err := r.TagsClient.GetAtScope(ctx, scope, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading subscription tags", fmt.Sprintf("Unable to read tags for subscription %s: %s", r.SubscriptionID, err))
		return
	}

	// Convert Azure tags to Terraform tags
	tfTags := make(map[string]string)
	if tagsResponse.Properties != nil && tagsResponse.Properties.Tags != nil {
		for k, v := range tagsResponse.Properties.Tags {
			if v != nil {
				tfTags[k] = *v
			}
		}
	}

	tagsValue, diags := types.MapValueFrom(ctx, types.StringType, tfTags)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Tags = tagsValue

	tagInheritance, err := r.SettingsClient.GetTagInheritance(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error getting tag inheritance settings", err.Error())
		return
	}

	if tagInheritance.Id != "" {
		data.InheritTags = types.BoolValue(true)
		data.PreferContainers = types.BoolValue(tagInheritance.Properties.PreferContainerTags)
	} else {
		data.InheritTags = types.BoolValue(false)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriptionTagsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SubscriptionTagsResourceModel
	var oldData *SubscriptionTagsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &oldData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "updating subscription tags resource")

	tfTags := make(map[string]string)
	resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &tfTags, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.applyTags(ctx, r.SubscriptionID, tfTags)
	if err != nil {
		resp.Diagnostics.AddError("Error updating subscription tags", err.Error())
		return
	}

	if data.InheritTags.ValueBool() {
		tagInheritance, err := r.SettingsClient.EnableTagInheritance(ctx, data.PreferContainers.ValueBool())
		if err != nil {
			resp.Diagnostics.AddError("Error updating tag inheritance settings", err.Error())
			return
		}

		if tagInheritance.Id != "" {
			data.InheritTags = types.BoolValue(true)
			data.PreferContainers = types.BoolValue(tagInheritance.Properties.PreferContainerTags)
		}
	} else if oldData.InheritTags.ValueBool() && !data.InheritTags.ValueBool() {
		_, err := r.SettingsClient.DisableTagInheritance(ctx)
		if err != nil {
			resp.Diagnostics.AddError("Error disabling tag inheritance", err.Error())
			return
		}
		data.InheritTags = types.BoolValue(false)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriptionTagsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SubscriptionTagsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "deleting subscription tags resource")

	if data.RemoveTags.ValueBool() {
		err := r.applyTags(ctx, r.SubscriptionID, map[string]string{})
		if err != nil {
			resp.Diagnostics.AddError("Error removing subscription tags", err.Error())
			return
		}
	}

	if data.RemoteInheritTags.ValueBool() && data.InheritTags.ValueBool() {
		_, err := r.SettingsClient.DisableTagInheritance(ctx)
		if err != nil {
			resp.Diagnostics.AddError("Error disabling tag inheritance", err.Error())
			return
		}
	}
}

func (r *SubscriptionTagsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("subscription_id"), req, resp)
}

func (r *SubscriptionTagsResource) applyTags(ctx context.Context, subscriptionID string, tagMap map[string]string) error {
	azureTags := make(map[string]*string)
	for k, v := range tagMap {
		value := v
		azureTags[k] = &value
	}

	scope := fmt.Sprintf("/subscriptions/%s", subscriptionID)

	_, err := r.TagsClient.CreateOrUpdateAtScope(ctx, scope, armresources.TagsResource{
		Properties: &armresources.Tags{
			Tags: azureTags,
		},
	}, nil)

	if err != nil {
		return fmt.Errorf("failed to set tags for subscription %q: %+v", subscriptionID, err)
	}

	return nil
}
