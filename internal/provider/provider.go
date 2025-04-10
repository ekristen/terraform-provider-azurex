// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/auth/autorest"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure AzurexProvider satisfies various provider interfaces.
var _ provider.Provider = &AzurexProvider{}
var _ provider.ProviderWithFunctions = &AzurexProvider{}
var _ provider.ProviderWithEphemeralResources = &AzurexProvider{}

// AzurexProvider defines the provider implementation.
type AzurexProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AzurexProviderModel describes the provider data model.
type AzurexProviderModel struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
	TenantID       types.String `tfsdk:"tenant_id"`
	ClientID       types.String `tfsdk:"client_id"`
	ClientSecret   types.String `tfsdk:"client_secret"`
}

type AzurexContext struct {
	SubscriptionID string

	Graph      *autorest.Authorizer
	Management *autorest.Authorizer

	MicrosoftGraph  auth.Authorizer
	ResourceManager auth.Authorizer

	IdentityCreds azcore.TokenCredential
}

func (p *AzurexProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "azurex"
	resp.Version = p.version
}

func (p *AzurexProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subscription_id": schema.StringAttribute{
				MarkdownDescription: "Azure Subscription ID",
				Optional:            true,
			},
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant ID",
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "SettingsClient ID",
				Optional:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "SettingsClient Secret",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *AzurexProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AzurexProviderModel
	var providerContext AzurexContext

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if v := os.Getenv("ARM_SUBSCRIPTION_ID"); v != "" {
		data.SubscriptionID = types.StringValue(v)
	}
	if v := os.Getenv("ARM_TENANT_ID"); v != "" {
		data.TenantID = types.StringValue(v)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	env, err := environments.FromName("global")
	if err != nil {
		resp.Diagnostics.AddError("unable to set environment", fmt.Sprintf("got: %s", err.Error()))
		return
	}

	credentials := auth.Credentials{
		Environment: *env,
		TenantID:    data.TenantID.ValueString(),
		ClientID:    os.Getenv("ARM_CLIENT_ID"),

		EnableAuthenticatingUsingClientSecret: true,
	}

	if os.Getenv("ARM_CLIENT_SECRET") != "" {
		tflog.Debug(ctx, "authentication type: client secret")
		credentials.EnableAuthenticatingUsingClientSecret = true
		credentials.ClientSecret = os.Getenv("ARM_CLIENT_SECRET")

		creds, err := azidentity.NewClientSecretCredential(data.TenantID.ValueString(), os.Getenv("ARM_CLIENT_ID"), os.Getenv("ARM_CLIENT_SECRET"), &azidentity.ClientSecretCredentialOptions{})
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}
		providerContext.IdentityCreds = creds
	} else if os.Getenv("ARM_CLIENT_CERTIFICATE_FILE") != "" {
		tflog.Debug(ctx, "authentication type: client certificate")
		credentials.EnableAuthenticatingUsingClientCertificate = true
		credentials.ClientCertificatePath = os.Getenv("ARM_CLIENT_CERTIFICATE_FILE")

		certData, err := os.ReadFile(os.Getenv("ARM_CLIENT_CERTIFICATE_FILE"))
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}

		certs, pkey, err := azidentity.ParseCertificates(certData, nil)
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}

		creds, err := azidentity.NewClientCertificateCredential(data.TenantID.ValueString(), os.Getenv("ARM_CLIENT_ID"), certs, pkey, &azidentity.ClientCertificateCredentialOptions{})
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}
		providerContext.IdentityCreds = creds
	} else if os.Getenv("AZURE_FEDERATED_TOKEN_FILE") != "" {
		tflog.Debug(ctx, "authentication type: federated token")
		token, err := os.ReadFile(os.Getenv("AZURE_FEDERATED_TOKEN_FILE"))
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}
		credentials.EnableAuthenticationUsingOIDC = true
		credentials.OIDCAssertionToken = string(token)

		creds, err := azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
			TenantID:      data.TenantID.ValueString(),
			ClientID:      os.Getenv("AZURE_CLIENT_ID"),
			TokenFilePath: os.Getenv("AZURE_FEDERATED_TOKEN_FILE"),
		})
		if err != nil {
			resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
			return
		}
		providerContext.IdentityCreds = creds
	}

	graphAuthorizer, err := auth.NewAuthorizerFromCredentials(ctx, credentials, env.MicrosoftGraph)
	if err != nil {
		resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
		return
	}

	mgmtAuthorizer, err := auth.NewAuthorizerFromCredentials(ctx, credentials, env.ResourceManager)
	if err != nil {
		resp.Diagnostics.AddError("unable to configure credential", fmt.Sprintf("got: %s", err.Error()))
		return
	}

	providerContext.Management = autorest.AutorestAuthorizer(mgmtAuthorizer)
	providerContext.Graph = autorest.AutorestAuthorizer(graphAuthorizer)

	providerContext.SubscriptionID = data.SubscriptionID.ValueString()

	resp.DataSourceData = providerContext
	resp.ResourceData = providerContext
}

func (p *AzurexProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSubscriptionTagsResource,
	}
}

func (p *AzurexProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *AzurexProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *AzurexProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AzurexProvider{
			version: version,
		}
	}
}
