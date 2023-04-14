// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum string

const (
	SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnumAccessToken SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum = "access_token"
)

func (e *SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "access_token":
		*e = SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum: %s", s)
	}
}

type SourceGitlabAuthorizationMethodPrivateToken struct {
	// Log into your Gitlab account and then generate a personal Access Token.
	AccessToken string                                                   `json:"access_token"`
	AuthType    *SourceGitlabAuthorizationMethodPrivateTokenAuthTypeEnum `json:"auth_type,omitempty"`
}

type SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum string

const (
	SourceGitlabAuthorizationMethodOAuth20AuthTypeEnumOauth20 SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum = "oauth2.0"
)

func (e *SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "oauth2.0":
		*e = SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum: %s", s)
	}
}

type SourceGitlabAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                              `json:"access_token"`
	AuthType    *SourceGitlabAuthorizationMethodOAuth20AuthTypeEnum `json:"auth_type,omitempty"`
	// The API ID of the Gitlab developer application.
	ClientID string `json:"client_id"`
	// The API Secret the Gitlab developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceGitlabAuthorizationMethodType string

const (
	SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodOAuth20      SourceGitlabAuthorizationMethodType = "source-gitlab_Authorization Method_OAuth2.0"
	SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodPrivateToken SourceGitlabAuthorizationMethodType = "source-gitlab_Authorization Method_Private Token"
)

type SourceGitlabAuthorizationMethod struct {
	SourceGitlabAuthorizationMethodOAuth20      *SourceGitlabAuthorizationMethodOAuth20
	SourceGitlabAuthorizationMethodPrivateToken *SourceGitlabAuthorizationMethodPrivateToken

	Type SourceGitlabAuthorizationMethodType
}

func CreateSourceGitlabAuthorizationMethodSourceGitlabAuthorizationMethodOAuth20(sourceGitlabAuthorizationMethodOAuth20 SourceGitlabAuthorizationMethodOAuth20) SourceGitlabAuthorizationMethod {
	typ := SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodOAuth20

	return SourceGitlabAuthorizationMethod{
		SourceGitlabAuthorizationMethodOAuth20: &sourceGitlabAuthorizationMethodOAuth20,
		Type:                                   typ,
	}
}

func CreateSourceGitlabAuthorizationMethodSourceGitlabAuthorizationMethodPrivateToken(sourceGitlabAuthorizationMethodPrivateToken SourceGitlabAuthorizationMethodPrivateToken) SourceGitlabAuthorizationMethod {
	typ := SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodPrivateToken

	return SourceGitlabAuthorizationMethod{
		SourceGitlabAuthorizationMethodPrivateToken: &sourceGitlabAuthorizationMethodPrivateToken,
		Type: typ,
	}
}

func (u *SourceGitlabAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceGitlabAuthorizationMethodOAuth20 := new(SourceGitlabAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGitlabAuthorizationMethodOAuth20); err == nil {
		u.SourceGitlabAuthorizationMethodOAuth20 = sourceGitlabAuthorizationMethodOAuth20
		u.Type = SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodOAuth20
		return nil
	}

	sourceGitlabAuthorizationMethodPrivateToken := new(SourceGitlabAuthorizationMethodPrivateToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGitlabAuthorizationMethodPrivateToken); err == nil {
		u.SourceGitlabAuthorizationMethodPrivateToken = sourceGitlabAuthorizationMethodPrivateToken
		u.Type = SourceGitlabAuthorizationMethodTypeSourceGitlabAuthorizationMethodPrivateToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGitlabAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceGitlabAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceGitlabAuthorizationMethodOAuth20)
	}

	if u.SourceGitlabAuthorizationMethodPrivateToken != nil {
		return json.Marshal(u.SourceGitlabAuthorizationMethodPrivateToken)
	}

	return nil, nil
}

type SourceGitlabGitlabEnum string

const (
	SourceGitlabGitlabEnumGitlab SourceGitlabGitlabEnum = "gitlab"
)

func (e *SourceGitlabGitlabEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "gitlab":
		*e = SourceGitlabGitlabEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGitlabGitlabEnum: %s", s)
	}
}

// SourceGitlab - The values required to configure the source.
type SourceGitlab struct {
	// Please enter your basic URL from GitLab instance.
	APIURL      string                          `json:"api_url"`
	Credentials SourceGitlabAuthorizationMethod `json:"credentials"`
	// Space-delimited list of groups. e.g. airbyte.io.
	Groups *string `json:"groups,omitempty"`
	// Space-delimited list of projects. e.g. airbyte.io/documentation meltano/tap-gitlab.
	Projects   *string                `json:"projects,omitempty"`
	SourceType SourceGitlabGitlabEnum `json:"sourceType"`
	// The date from which you'd like to replicate data for GitLab API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
}