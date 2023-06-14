// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceFreshserviceResourceModel) ToCreateSDKType() *shared.SourceFreshserviceCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	domainName := r.Configuration.DomainName.ValueString()
	sourceType := shared.SourceFreshserviceFreshservice(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceFreshservice{
		APIKey:     apiKey,
		DomainName: domainName,
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
	out := shared.SourceFreshserviceCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFreshserviceResourceModel) ToGetSDKType() *shared.SourceFreshserviceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFreshserviceResourceModel) ToUpdateSDKType() *shared.SourceFreshservicePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	domainName := r.Configuration.DomainName.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceFreshserviceUpdate{
		APIKey:     apiKey,
		DomainName: domainName,
		StartDate:  startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFreshservicePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFreshserviceResourceModel) ToDeleteSDKType() *shared.SourceFreshserviceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFreshserviceResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFreshserviceResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}