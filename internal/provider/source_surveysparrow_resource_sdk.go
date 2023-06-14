// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSurveySparrowResourceModel) ToCreateSDKType() *shared.SourceSurveySparrowCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	var region *shared.SourceSurveySparrowBaseURL
	var sourceSurveySparrowBaseURLEUBasedAccount *shared.SourceSurveySparrowBaseURLEUBasedAccount
	if r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount != nil {
		urlBase := new(shared.SourceSurveySparrowBaseURLEUBasedAccountURLBase)
		if !r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.IsUnknown() && !r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.IsNull() {
			*urlBase = shared.SourceSurveySparrowBaseURLEUBasedAccountURLBase(r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.ValueString())
		} else {
			urlBase = nil
		}
		sourceSurveySparrowBaseURLEUBasedAccount = &shared.SourceSurveySparrowBaseURLEUBasedAccount{
			URLBase: urlBase,
		}
	}
	if sourceSurveySparrowBaseURLEUBasedAccount != nil {
		region = &shared.SourceSurveySparrowBaseURL{
			SourceSurveySparrowBaseURLEUBasedAccount: sourceSurveySparrowBaseURLEUBasedAccount,
		}
	}
	var sourceSurveySparrowBaseURLGlobalAccount *shared.SourceSurveySparrowBaseURLGlobalAccount
	if r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount != nil {
		urlBase1 := new(shared.SourceSurveySparrowBaseURLGlobalAccountURLBase)
		if !r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.IsUnknown() && !r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.IsNull() {
			*urlBase1 = shared.SourceSurveySparrowBaseURLGlobalAccountURLBase(r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.ValueString())
		} else {
			urlBase1 = nil
		}
		sourceSurveySparrowBaseURLGlobalAccount = &shared.SourceSurveySparrowBaseURLGlobalAccount{
			URLBase: urlBase1,
		}
	}
	if sourceSurveySparrowBaseURLGlobalAccount != nil {
		region = &shared.SourceSurveySparrowBaseURL{
			SourceSurveySparrowBaseURLGlobalAccount: sourceSurveySparrowBaseURLGlobalAccount,
		}
	}
	sourceType := shared.SourceSurveySparrowSurveySparrow(r.Configuration.SourceType.ValueString())
	surveyID := make([]interface{}, 0)
	for _, surveyIDItem := range r.Configuration.SurveyID {
		var surveyIDTmp interface{}
		_ = json.Unmarshal([]byte(surveyIDItem.ValueString()), &surveyIDTmp)
		surveyID = append(surveyID, surveyIDTmp)
	}
	configuration := shared.SourceSurveySparrow{
		AccessToken: accessToken,
		Region:      region,
		SourceType:  sourceType,
		SurveyID:    surveyID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSurveySparrowCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSurveySparrowResourceModel) ToGetSDKType() *shared.SourceSurveySparrowCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSurveySparrowResourceModel) ToUpdateSDKType() *shared.SourceSurveySparrowPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	var region *shared.SourceSurveySparrowUpdateBaseURL
	var sourceSurveySparrowUpdateBaseURLEUBasedAccount *shared.SourceSurveySparrowUpdateBaseURLEUBasedAccount
	if r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount != nil {
		urlBase := new(shared.SourceSurveySparrowUpdateBaseURLEUBasedAccountURLBase)
		if !r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.IsUnknown() && !r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.IsNull() {
			*urlBase = shared.SourceSurveySparrowUpdateBaseURLEUBasedAccountURLBase(r.Configuration.Region.SourceSurveySparrowBaseURLEUBasedAccount.URLBase.ValueString())
		} else {
			urlBase = nil
		}
		sourceSurveySparrowUpdateBaseURLEUBasedAccount = &shared.SourceSurveySparrowUpdateBaseURLEUBasedAccount{
			URLBase: urlBase,
		}
	}
	if sourceSurveySparrowUpdateBaseURLEUBasedAccount != nil {
		region = &shared.SourceSurveySparrowUpdateBaseURL{
			SourceSurveySparrowUpdateBaseURLEUBasedAccount: sourceSurveySparrowUpdateBaseURLEUBasedAccount,
		}
	}
	var sourceSurveySparrowUpdateBaseURLGlobalAccount *shared.SourceSurveySparrowUpdateBaseURLGlobalAccount
	if r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount != nil {
		urlBase1 := new(shared.SourceSurveySparrowUpdateBaseURLGlobalAccountURLBase)
		if !r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.IsUnknown() && !r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.IsNull() {
			*urlBase1 = shared.SourceSurveySparrowUpdateBaseURLGlobalAccountURLBase(r.Configuration.Region.SourceSurveySparrowBaseURLGlobalAccount.URLBase.ValueString())
		} else {
			urlBase1 = nil
		}
		sourceSurveySparrowUpdateBaseURLGlobalAccount = &shared.SourceSurveySparrowUpdateBaseURLGlobalAccount{
			URLBase: urlBase1,
		}
	}
	if sourceSurveySparrowUpdateBaseURLGlobalAccount != nil {
		region = &shared.SourceSurveySparrowUpdateBaseURL{
			SourceSurveySparrowUpdateBaseURLGlobalAccount: sourceSurveySparrowUpdateBaseURLGlobalAccount,
		}
	}
	surveyID := make([]interface{}, 0)
	for _, surveyIDItem := range r.Configuration.SurveyID {
		var surveyIDTmp interface{}
		_ = json.Unmarshal([]byte(surveyIDItem.ValueString()), &surveyIDTmp)
		surveyID = append(surveyID, surveyIDTmp)
	}
	configuration := shared.SourceSurveySparrowUpdate{
		AccessToken: accessToken,
		Region:      region,
		SurveyID:    surveyID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSurveySparrowPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSurveySparrowResourceModel) ToDeleteSDKType() *shared.SourceSurveySparrowCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSurveySparrowResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSurveySparrowResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}