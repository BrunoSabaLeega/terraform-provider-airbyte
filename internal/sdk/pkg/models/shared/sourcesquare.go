// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"openapi/internal/sdk/pkg/types"
)

type SourceSquareAuthenticationAPIKeyCredentialsTitleEnum string

const (
	SourceSquareAuthenticationAPIKeyCredentialsTitleEnumAPIKey SourceSquareAuthenticationAPIKeyCredentialsTitleEnum = "API Key"
)

func (e *SourceSquareAuthenticationAPIKeyCredentialsTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "API Key":
		*e = SourceSquareAuthenticationAPIKeyCredentialsTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareAuthenticationAPIKeyCredentialsTitleEnum: %s", s)
	}
}

// SourceSquareAuthenticationAPIKey - Choose how to authenticate to Square.
type SourceSquareAuthenticationAPIKey struct {
	// The API key for a Square application
	APIKey           string                                                `json:"api_key"`
	CredentialsTitle *SourceSquareAuthenticationAPIKeyCredentialsTitleEnum `json:"credentials_title,omitempty"`
}

type SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum string

const (
	SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnumOAuthCredentials SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum = "OAuth Credentials"
)

func (e *SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "OAuth Credentials":
		*e = SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum: %s", s)
	}
}

// SourceSquareAuthenticationOauthAuthentication - Choose how to authenticate to Square.
type SourceSquareAuthenticationOauthAuthentication struct {
	// The Square-issued ID of your application
	ClientID string `json:"client_id"`
	// The Square-issued application secret for your application
	ClientSecret     string                                                             `json:"client_secret"`
	CredentialsTitle *SourceSquareAuthenticationOauthAuthenticationCredentialsTitleEnum `json:"credentials_title,omitempty"`
	// A refresh token generated using the above client ID and secret
	RefreshToken string `json:"refresh_token"`
}

type SourceSquareAuthenticationType string

const (
	SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication SourceSquareAuthenticationType = "source-square_Authentication_Oauth authentication"
	SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey              SourceSquareAuthenticationType = "source-square_Authentication_API key"
)

type SourceSquareAuthentication struct {
	SourceSquareAuthenticationOauthAuthentication *SourceSquareAuthenticationOauthAuthentication
	SourceSquareAuthenticationAPIKey              *SourceSquareAuthenticationAPIKey

	Type SourceSquareAuthenticationType
}

func CreateSourceSquareAuthenticationSourceSquareAuthenticationOauthAuthentication(sourceSquareAuthenticationOauthAuthentication SourceSquareAuthenticationOauthAuthentication) SourceSquareAuthentication {
	typ := SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication

	return SourceSquareAuthentication{
		SourceSquareAuthenticationOauthAuthentication: &sourceSquareAuthenticationOauthAuthentication,
		Type: typ,
	}
}

func CreateSourceSquareAuthenticationSourceSquareAuthenticationAPIKey(sourceSquareAuthenticationAPIKey SourceSquareAuthenticationAPIKey) SourceSquareAuthentication {
	typ := SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey

	return SourceSquareAuthentication{
		SourceSquareAuthenticationAPIKey: &sourceSquareAuthenticationAPIKey,
		Type:                             typ,
	}
}

func (u *SourceSquareAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSquareAuthenticationOauthAuthentication := new(SourceSquareAuthenticationOauthAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSquareAuthenticationOauthAuthentication); err == nil {
		u.SourceSquareAuthenticationOauthAuthentication = sourceSquareAuthenticationOauthAuthentication
		u.Type = SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication
		return nil
	}

	sourceSquareAuthenticationAPIKey := new(SourceSquareAuthenticationAPIKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSquareAuthenticationAPIKey); err == nil {
		u.SourceSquareAuthenticationAPIKey = sourceSquareAuthenticationAPIKey
		u.Type = SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSquareAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceSquareAuthenticationOauthAuthentication != nil {
		return json.Marshal(u.SourceSquareAuthenticationOauthAuthentication)
	}

	if u.SourceSquareAuthenticationAPIKey != nil {
		return json.Marshal(u.SourceSquareAuthenticationAPIKey)
	}

	return nil, nil
}

type SourceSquareSquareEnum string

const (
	SourceSquareSquareEnumSquare SourceSquareSquareEnum = "square"
)

func (e *SourceSquareSquareEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "square":
		*e = SourceSquareSquareEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareSquareEnum: %s", s)
	}
}

// SourceSquare - The values required to configure the source.
type SourceSquare struct {
	// Choose how to authenticate to Square.
	Credentials *SourceSquareAuthentication `json:"credentials,omitempty"`
	// In some streams there is an option to include deleted objects (Items, Categories, Discounts, Taxes)
	IncludeDeletedObjects *bool `json:"include_deleted_objects,omitempty"`
	// Determines whether to use the sandbox or production environment.
	IsSandbox  bool                   `json:"is_sandbox"`
	SourceType SourceSquareSquareEnum `json:"sourceType"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. If not set, all data will be replicated.
	StartDate *types.Date `json:"start_date,omitempty"`
}