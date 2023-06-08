// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePinterestResourceModel) ToCreateSDKType() *shared.SourcePinterestCreateRequest {
	var credentials *shared.SourcePinterestAuthorizationMethod
	var sourcePinterestAuthorizationMethodOAuth20 *shared.SourcePinterestAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourcePinterestAuthorizationMethodAccessToken != nil {
		authMethod := shared.SourcePinterestAuthorizationMethodOAuth20AuthMethod(r.Configuration.Credentials.SourcePinterestAuthorizationMethodAccessToken.AuthMethod.ValueString())
		sourcePinterestAuthorizationMethodOAuth20 = &shared.SourcePinterestAuthorizationMethodOAuth20{
			AuthMethod: authMethod,
		}
	}
	if sourcePinterestAuthorizationMethodOAuth20 != nil {
		credentials = &shared.SourcePinterestAuthorizationMethod{
			SourcePinterestAuthorizationMethodOAuth20: sourcePinterestAuthorizationMethodOAuth20,
		}
	}
	var sourcePinterestAuthorizationMethodAccessToken *shared.SourcePinterestAuthorizationMethodAccessToken
	if r.Configuration.Credentials.SourcePinterestAuthorizationMethodOAuth20 != nil {
		authMethod1 := shared.SourcePinterestAuthorizationMethodAccessTokenAuthMethod(r.Configuration.Credentials.SourcePinterestAuthorizationMethodOAuth20.AuthMethod.ValueString())
		sourcePinterestAuthorizationMethodAccessToken = &shared.SourcePinterestAuthorizationMethodAccessToken{
			AuthMethod: authMethod1,
		}
	}
	if sourcePinterestAuthorizationMethodAccessToken != nil {
		credentials = &shared.SourcePinterestAuthorizationMethod{
			SourcePinterestAuthorizationMethodAccessToken: sourcePinterestAuthorizationMethodAccessToken,
		}
	}
	sourceType := shared.SourcePinterestPinterest(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	status := make([]shared.SourcePinterestStatus, 0)
	for _, statusItem := range r.Configuration.Status {
		status = append(status, shared.SourcePinterestStatus(statusItem.ValueString()))
	}
	configuration := shared.SourcePinterest{
		Credentials: credentials,
		SourceType:  sourceType,
		StartDate:   startDate,
		Status:      status,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePinterestCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePinterestResourceModel) ToUpdateSDKType() *shared.SourcePinterestPutRequest {
	var credentials *shared.SourcePinterestUpdateAuthorizationMethod
	var sourcePinterestUpdateAuthorizationMethodOAuth20 *shared.SourcePinterestUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourcePinterestAuthorizationMethodAccessToken != nil {
		authMethod := shared.SourcePinterestUpdateAuthorizationMethodOAuth20AuthMethod(r.Configuration.Credentials.SourcePinterestAuthorizationMethodAccessToken.AuthMethod.ValueString())
		sourcePinterestUpdateAuthorizationMethodOAuth20 = &shared.SourcePinterestUpdateAuthorizationMethodOAuth20{
			AuthMethod: authMethod,
		}
	}
	if sourcePinterestUpdateAuthorizationMethodOAuth20 != nil {
		credentials = &shared.SourcePinterestUpdateAuthorizationMethod{
			SourcePinterestUpdateAuthorizationMethodOAuth20: sourcePinterestUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourcePinterestUpdateAuthorizationMethodAccessToken *shared.SourcePinterestUpdateAuthorizationMethodAccessToken
	if r.Configuration.Credentials.SourcePinterestAuthorizationMethodOAuth20 != nil {
		authMethod1 := shared.SourcePinterestUpdateAuthorizationMethodAccessTokenAuthMethod(r.Configuration.Credentials.SourcePinterestAuthorizationMethodOAuth20.AuthMethod.ValueString())
		sourcePinterestUpdateAuthorizationMethodAccessToken = &shared.SourcePinterestUpdateAuthorizationMethodAccessToken{
			AuthMethod: authMethod1,
		}
	}
	if sourcePinterestUpdateAuthorizationMethodAccessToken != nil {
		credentials = &shared.SourcePinterestUpdateAuthorizationMethod{
			SourcePinterestUpdateAuthorizationMethodAccessToken: sourcePinterestUpdateAuthorizationMethodAccessToken,
		}
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	status := make([]shared.SourcePinterestUpdateStatus, 0)
	for _, statusItem := range r.Configuration.Status {
		status = append(status, shared.SourcePinterestUpdateStatus(statusItem.ValueString()))
	}
	configuration := shared.SourcePinterestUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Status:      status,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePinterestPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePinterestResourceModel) ToDeleteSDKType() *shared.SourcePinterestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePinterestResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
