// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceK6CloudResourceModel) ToCreateSDKType() *shared.SourceK6CloudCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	sourceType := shared.SourceK6CloudK6Cloud(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceK6Cloud{
		APIToken:   apiToken,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceK6CloudCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceK6CloudResourceModel) ToGetSDKType() *shared.SourceK6CloudCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceK6CloudResourceModel) ToUpdateSDKType() *shared.SourceK6CloudPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	configuration := shared.SourceK6CloudUpdate{
		APIToken: apiToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceK6CloudPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceK6CloudResourceModel) ToDeleteSDKType() *shared.SourceK6CloudCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceK6CloudResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceK6CloudResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}