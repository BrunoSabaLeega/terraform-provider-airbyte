// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceKlaviyoResourceModel) ToCreateSDKType() *shared.SourceKlaviyoCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceKlaviyoKlaviyo(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceKlaviyo{
		APIKey:     apiKey,
		SourceType: sourceType,
		StartDate:  startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKlaviyoCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKlaviyoResourceModel) ToGetSDKType() *shared.SourceKlaviyoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKlaviyoResourceModel) ToUpdateSDKType() *shared.SourceKlaviyoPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceKlaviyoUpdate{
		APIKey:    apiKey,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKlaviyoPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKlaviyoResourceModel) ToDeleteSDKType() *shared.SourceKlaviyoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKlaviyoResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceKlaviyoResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}