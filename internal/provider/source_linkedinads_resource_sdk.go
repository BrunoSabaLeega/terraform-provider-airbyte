// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLinkedinAdsResourceModel) ToCreateSDKType() *shared.SourceLinkedinAdsCreateRequest {
	accountIds := make([]int64, 0)
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var credentials *shared.SourceLinkedinAdsAuthentication
	var sourceLinkedinAdsAuthenticationOAuth20 *shared.SourceLinkedinAdsAuthenticationOAuth20
	if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken != nil {
		authMethod := new(shared.SourceLinkedinAdsAuthenticationOAuth20AuthMethod)
		if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsNull() {
			*authMethod = shared.SourceLinkedinAdsAuthenticationOAuth20AuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.ValueString())
		} else {
			authMethod = nil
		}
		sourceLinkedinAdsAuthenticationOAuth20 = &shared.SourceLinkedinAdsAuthenticationOAuth20{
			AuthMethod: authMethod,
		}
	}
	if sourceLinkedinAdsAuthenticationOAuth20 != nil {
		credentials = &shared.SourceLinkedinAdsAuthentication{
			SourceLinkedinAdsAuthenticationOAuth20: sourceLinkedinAdsAuthenticationOAuth20,
		}
	}
	var sourceLinkedinAdsAuthenticationAccessToken *shared.SourceLinkedinAdsAuthenticationAccessToken
	if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20 != nil {
		authMethod1 := new(shared.SourceLinkedinAdsAuthenticationAccessTokenAuthMethod)
		if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsNull() {
			*authMethod1 = shared.SourceLinkedinAdsAuthenticationAccessTokenAuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.ValueString())
		} else {
			authMethod1 = nil
		}
		sourceLinkedinAdsAuthenticationAccessToken = &shared.SourceLinkedinAdsAuthenticationAccessToken{
			AuthMethod: authMethod1,
		}
	}
	if sourceLinkedinAdsAuthenticationAccessToken != nil {
		credentials = &shared.SourceLinkedinAdsAuthentication{
			SourceLinkedinAdsAuthenticationAccessToken: sourceLinkedinAdsAuthenticationAccessToken,
		}
	}
	sourceType := shared.SourceLinkedinAdsLinkedinAds(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAds{
		AccountIds:  accountIds,
		Credentials: credentials,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinAdsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinAdsResourceModel) ToGetSDKType() *shared.SourceLinkedinAdsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinAdsResourceModel) ToUpdateSDKType() *shared.SourceLinkedinAdsPutRequest {
	accountIds := make([]int64, 0)
	for _, accountIdsItem := range r.Configuration.AccountIds {
		accountIds = append(accountIds, accountIdsItem.ValueInt64())
	}
	var credentials *shared.SourceLinkedinAdsUpdateAuthentication
	var sourceLinkedinAdsUpdateAuthenticationOAuth20 *shared.SourceLinkedinAdsUpdateAuthenticationOAuth20
	if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken != nil {
		authMethod := new(shared.SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod)
		if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.IsNull() {
			*authMethod = shared.SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationAccessToken.AuthMethod.ValueString())
		} else {
			authMethod = nil
		}
		sourceLinkedinAdsUpdateAuthenticationOAuth20 = &shared.SourceLinkedinAdsUpdateAuthenticationOAuth20{
			AuthMethod: authMethod,
		}
	}
	if sourceLinkedinAdsUpdateAuthenticationOAuth20 != nil {
		credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
			SourceLinkedinAdsUpdateAuthenticationOAuth20: sourceLinkedinAdsUpdateAuthenticationOAuth20,
		}
	}
	var sourceLinkedinAdsUpdateAuthenticationAccessToken *shared.SourceLinkedinAdsUpdateAuthenticationAccessToken
	if r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20 != nil {
		authMethod1 := new(shared.SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod)
		if !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsUnknown() && !r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.IsNull() {
			*authMethod1 = shared.SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod(r.Configuration.Credentials.SourceLinkedinAdsAuthenticationOAuth20.AuthMethod.ValueString())
		} else {
			authMethod1 = nil
		}
		sourceLinkedinAdsUpdateAuthenticationAccessToken = &shared.SourceLinkedinAdsUpdateAuthenticationAccessToken{
			AuthMethod: authMethod1,
		}
	}
	if sourceLinkedinAdsUpdateAuthenticationAccessToken != nil {
		credentials = &shared.SourceLinkedinAdsUpdateAuthentication{
			SourceLinkedinAdsUpdateAuthenticationAccessToken: sourceLinkedinAdsUpdateAuthenticationAccessToken,
		}
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceLinkedinAdsUpdate{
		AccountIds:  accountIds,
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinAdsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinAdsResourceModel) ToDeleteSDKType() *shared.SourceLinkedinAdsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinAdsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceLinkedinAdsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}