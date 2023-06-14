// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleWebfontsResourceModel) ToCreateSDKType() *shared.SourceGoogleWebfontsCreateRequest {
	alt := new(string)
	if !r.Configuration.Alt.IsUnknown() && !r.Configuration.Alt.IsNull() {
		*alt = r.Configuration.Alt.ValueString()
	} else {
		alt = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	prettyPrint := new(string)
	if !r.Configuration.PrettyPrint.IsUnknown() && !r.Configuration.PrettyPrint.IsNull() {
		*prettyPrint = r.Configuration.PrettyPrint.ValueString()
	} else {
		prettyPrint = nil
	}
	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	sourceType := shared.SourceGoogleWebfontsGoogleWebfonts(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceGoogleWebfonts{
		Alt:         alt,
		APIKey:      apiKey,
		PrettyPrint: prettyPrint,
		Sort:        sort,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleWebfontsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleWebfontsResourceModel) ToGetSDKType() *shared.SourceGoogleWebfontsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleWebfontsResourceModel) ToUpdateSDKType() *shared.SourceGoogleWebfontsPutRequest {
	alt := new(string)
	if !r.Configuration.Alt.IsUnknown() && !r.Configuration.Alt.IsNull() {
		*alt = r.Configuration.Alt.ValueString()
	} else {
		alt = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	prettyPrint := new(string)
	if !r.Configuration.PrettyPrint.IsUnknown() && !r.Configuration.PrettyPrint.IsNull() {
		*prettyPrint = r.Configuration.PrettyPrint.ValueString()
	} else {
		prettyPrint = nil
	}
	sort := new(string)
	if !r.Configuration.Sort.IsUnknown() && !r.Configuration.Sort.IsNull() {
		*sort = r.Configuration.Sort.ValueString()
	} else {
		sort = nil
	}
	configuration := shared.SourceGoogleWebfontsUpdate{
		Alt:         alt,
		APIKey:      apiKey,
		PrettyPrint: prettyPrint,
		Sort:        sort,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleWebfontsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleWebfontsResourceModel) ToDeleteSDKType() *shared.SourceGoogleWebfontsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleWebfontsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGoogleWebfontsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}