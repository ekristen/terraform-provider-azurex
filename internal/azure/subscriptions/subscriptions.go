// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptions

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SettingsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type SettingsClient struct {
	internal *arm.Client
}

// NewSettingsClient creates a new instance of SettingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".SettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SettingsClient{
		internal: cl,
	}
	return client, nil
}

type TagInheritanceProperties struct {
	PreferContainerTags bool `json:"preferContainerTags"`
}

// TagInheritanceRequest
// {"kind":"taginheritance","properties":{"preferContainerTags":false}}
type TagInheritanceRequest struct {
	Kind       string                   `json:"kind"`
	Properties TagInheritanceProperties `json:"properties"`
}

// TagInheritanceResponse
// {"properties":{"preferContainerTags":false},"id":"/subscriptions/a4c52fbc-96a6-43f5-b093-2188b94952a6/providers/Microsoft.CostManagement/settings/taginheritance","name":"taginheritance","type":"Microsoft.CostManagement/Settings","scope":"Subscription"}
type TagInheritanceResponse struct {
	Id         string                   `json:"id"`
	Name       string                   `json:"name"`
	Type       string                   `json:"type"`
	Scope      string                   `json:"scope"`
	Properties TagInheritanceProperties `json:"properties"`
}

func (client *SettingsClient) GetTagInheritance(ctx context.Context) (TagInheritanceResponse, error) {
	req, err := client.getTagInheritanceRequest(ctx)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	return client.handleTagInheritanceResponse(resp)
}

func (client *SettingsClient) EnableTagInheritance(ctx context.Context, preferContainerTags bool) (TagInheritanceResponse, error) {
	req, err := client.createTagInheritanceRequest(ctx, preferContainerTags)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	return client.handleTagInheritanceResponse(resp)
}

func (client *SettingsClient) DisableTagInheritance(ctx context.Context) (TagInheritanceResponse, error) {
	req, err := client.createTagInheritanceRequest(ctx, false)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagInheritanceResponse{}, err
	}

	return client.handleTagInheritanceResponse(resp)
}

func (client *SettingsClient) getTagInheritanceRequest(ctx context.Context) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/settings/taginheritance"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createTagInheritanceRequest
// https://management.azure.com/`subscription`s/a4c52fbc-96a6-43f5-b093-2188b94952a6/providers/Microsoft.CostManagement/settings/taginheritance?api-version=2022-10-01-preview
func (client *SettingsClient) createTagInheritanceRequest(ctx context.Context, preferContainerTags bool) (*policy.Request, error) {
	params := TagInheritanceRequest{
		Kind: "taginheritance",
		Properties: TagInheritanceProperties{
			PreferContainerTags: preferContainerTags,
		},
	}
	urlPath := "/providers/Microsoft.CostManagement/settings/taginheritance"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, params)
}

func (client *SettingsClient) handleTagInheritanceResponse(resp *http.Response) (TagInheritanceResponse, error) {
	var TagInheritance TagInheritanceResponse

	if err := runtime.UnmarshalAsJSON(resp, &TagInheritance); err != nil {
		return TagInheritanceResponse{}, err
	}
	return TagInheritance, nil
}
