// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum string

const (
	SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnumAccessToken SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum = "access_token"
)

func (e *SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "access_token":
		*e = SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum: %s", s)
	}
}

type SourceZendeskChatAuthorizationMethodAccessToken struct {
	// The Access Token to make authenticated requests.
	AccessToken string                                                         `json:"access_token"`
	Credentials SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum `json:"credentials"`
}

type SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum string

const (
	SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnumOauth20 SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum = "oauth2.0"
)

func (e *SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "oauth2.0":
		*e = SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum: %s", s)
	}
}

type SourceZendeskChatAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string `json:"access_token,omitempty"`
	// The Client ID of your OAuth application
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string                                                    `json:"client_secret,omitempty"`
	Credentials  SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum `json:"credentials"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type SourceZendeskChatAuthorizationMethodType string

const (
	SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20     SourceZendeskChatAuthorizationMethodType = "source-zendesk-chat_Authorization Method_OAuth2.0"
	SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken SourceZendeskChatAuthorizationMethodType = "source-zendesk-chat_Authorization Method_Access Token"
)

type SourceZendeskChatAuthorizationMethod struct {
	SourceZendeskChatAuthorizationMethodOAuth20     *SourceZendeskChatAuthorizationMethodOAuth20
	SourceZendeskChatAuthorizationMethodAccessToken *SourceZendeskChatAuthorizationMethodAccessToken

	Type SourceZendeskChatAuthorizationMethodType
}

func CreateSourceZendeskChatAuthorizationMethodSourceZendeskChatAuthorizationMethodOAuth20(sourceZendeskChatAuthorizationMethodOAuth20 SourceZendeskChatAuthorizationMethodOAuth20) SourceZendeskChatAuthorizationMethod {
	typ := SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20

	return SourceZendeskChatAuthorizationMethod{
		SourceZendeskChatAuthorizationMethodOAuth20: &sourceZendeskChatAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceZendeskChatAuthorizationMethodSourceZendeskChatAuthorizationMethodAccessToken(sourceZendeskChatAuthorizationMethodAccessToken SourceZendeskChatAuthorizationMethodAccessToken) SourceZendeskChatAuthorizationMethod {
	typ := SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken

	return SourceZendeskChatAuthorizationMethod{
		SourceZendeskChatAuthorizationMethodAccessToken: &sourceZendeskChatAuthorizationMethodAccessToken,
		Type: typ,
	}
}

func (u *SourceZendeskChatAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceZendeskChatAuthorizationMethodOAuth20 := new(SourceZendeskChatAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatAuthorizationMethodOAuth20); err == nil {
		u.SourceZendeskChatAuthorizationMethodOAuth20 = sourceZendeskChatAuthorizationMethodOAuth20
		u.Type = SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20
		return nil
	}

	sourceZendeskChatAuthorizationMethodAccessToken := new(SourceZendeskChatAuthorizationMethodAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatAuthorizationMethodAccessToken); err == nil {
		u.SourceZendeskChatAuthorizationMethodAccessToken = sourceZendeskChatAuthorizationMethodAccessToken
		u.Type = SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskChatAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskChatAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceZendeskChatAuthorizationMethodOAuth20)
	}

	if u.SourceZendeskChatAuthorizationMethodAccessToken != nil {
		return json.Marshal(u.SourceZendeskChatAuthorizationMethodAccessToken)
	}

	return nil, nil
}

type SourceZendeskChatZendeskChatEnum string

const (
	SourceZendeskChatZendeskChatEnumZendeskChat SourceZendeskChatZendeskChatEnum = "zendesk-chat"
)

func (e *SourceZendeskChatZendeskChatEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "zendesk-chat":
		*e = SourceZendeskChatZendeskChatEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatZendeskChatEnum: %s", s)
	}
}

// SourceZendeskChat - The values required to configure the source.
type SourceZendeskChat struct {
	Credentials *SourceZendeskChatAuthorizationMethod `json:"credentials,omitempty"`
	SourceType  SourceZendeskChatZendeskChatEnum      `json:"sourceType"`
	// The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// Required if you access Zendesk Chat from a Zendesk Support subdomain.
	Subdomain *string `json:"subdomain,omitempty"`
}