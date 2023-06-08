// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleAnalyticsDataAPIResourceModel) ToCreateSDKType() *shared.SourceGoogleAnalyticsDataAPICreateRequest {
	var credentials *shared.SourceGoogleAnalyticsDataAPICredentials
	var sourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth *shared.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth
	if r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth != nil {
		accessToken := new(string)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.IsNull() {
			*accessToken = r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		authType := new(shared.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauthAuthType)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.IsNull() {
			*authType = shared.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauthAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.RefreshToken.ValueString()
		sourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth = &shared.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth != nil {
		credentials = &shared.SourceGoogleAnalyticsDataAPICredentials{
			SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth: sourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth,
		}
	}
	var sourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication
	if r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication != nil {
		authType1 := new(shared.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthenticationAuthType)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.IsNull() {
			*authType1 = shared.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		credentialsJSON := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
		sourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication{
			AuthType:        authType1,
			CredentialsJSON: credentialsJSON,
		}
	}
	if sourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication != nil {
		credentials = &shared.SourceGoogleAnalyticsDataAPICredentials{
			SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication: sourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication,
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	dateRangesStartDate := customTypes.MustDateFromString(r.Configuration.DateRangesStartDate.ValueString())
	propertyID := r.Configuration.PropertyID.ValueString()
	sourceType := shared.SourceGoogleAnalyticsDataAPIGoogleAnalyticsDataAPI(r.Configuration.SourceType.ValueString())
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsDataAPI{
		Credentials:         credentials,
		CustomReports:       customReports,
		DateRangesStartDate: dateRangesStartDate,
		PropertyID:          propertyID,
		SourceType:          sourceType,
		WindowInDays:        windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsDataAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsDataAPIResourceModel) ToUpdateSDKType() *shared.SourceGoogleAnalyticsDataAPIPutRequest {
	var credentials *shared.SourceGoogleAnalyticsDataAPIUpdateCredentials
	var sourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth *shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth
	if r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth != nil {
		accessToken := new(string)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.IsNull() {
			*accessToken = r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		authType := new(shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauthAuthType)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.IsNull() {
			*authType = shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauthAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsAuthenticateViaGoogleOauth.RefreshToken.ValueString()
		sourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth = &shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth != nil {
		credentials = &shared.SourceGoogleAnalyticsDataAPIUpdateCredentials{
			SourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth: sourceGoogleAnalyticsDataAPIUpdateCredentialsAuthenticateViaGoogleOauth,
		}
	}
	var sourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication
	if r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication != nil {
		authType1 := new(shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthenticationAuthType)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.IsNull() {
			*authType1 = shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		credentialsJSON := r.Configuration.Credentials.SourceGoogleAnalyticsDataAPICredentialsServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
		sourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication{
			AuthType:        authType1,
			CredentialsJSON: credentialsJSON,
		}
	}
	if sourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication != nil {
		credentials = &shared.SourceGoogleAnalyticsDataAPIUpdateCredentials{
			SourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication: sourceGoogleAnalyticsDataAPIUpdateCredentialsServiceAccountKeyAuthentication,
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	dateRangesStartDate := customTypes.MustDateFromString(r.Configuration.DateRangesStartDate.ValueString())
	propertyID := r.Configuration.PropertyID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsDataAPIUpdate{
		Credentials:         credentials,
		CustomReports:       customReports,
		DateRangesStartDate: dateRangesStartDate,
		PropertyID:          propertyID,
		WindowInDays:        windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsDataAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsDataAPIResourceModel) ToDeleteSDKType() *shared.SourceGoogleAnalyticsDataAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleAnalyticsDataAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
