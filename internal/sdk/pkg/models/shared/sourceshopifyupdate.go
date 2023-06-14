// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod string

const (
	SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethodOauth20 SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod = "oauth2.0"
)

func (e SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod) ToPointer() *SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod {
	return &e
}

func (e *SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod: %v", v)
	}
}

// SourceShopifyUpdateShopifyAuthorizationMethodOAuth20 - OAuth2.0
type SourceShopifyUpdateShopifyAuthorizationMethodOAuth20 struct {
	// The Access Token for making authenticated requests.
	AccessToken *string                                                        `json:"access_token,omitempty"`
	AuthMethod  SourceShopifyUpdateShopifyAuthorizationMethodOAuth20AuthMethod `json:"auth_method"`
	// The Client ID of the Shopify developer application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of the Shopify developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
}

type SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod string

const (
	SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethodAPIPassword SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod = "api_password"
)

func (e SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod) ToPointer() *SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod {
	return &e
}

func (e *SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_password":
		*e = SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod: %v", v)
	}
}

// SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword - API Password Auth
type SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword struct {
	// The API Password for your private application in the `Shopify` store.
	APIPassword string                                                             `json:"api_password"`
	AuthMethod  SourceShopifyUpdateShopifyAuthorizationMethodAPIPasswordAuthMethod `json:"auth_method"`
}

type SourceShopifyUpdateShopifyAuthorizationMethodType string

const (
	SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodAPIPassword SourceShopifyUpdateShopifyAuthorizationMethodType = "source-shopify-update_Shopify Authorization Method_API Password"
	SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodOAuth20     SourceShopifyUpdateShopifyAuthorizationMethodType = "source-shopify-update_Shopify Authorization Method_OAuth2.0"
)

type SourceShopifyUpdateShopifyAuthorizationMethod struct {
	SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword *SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword
	SourceShopifyUpdateShopifyAuthorizationMethodOAuth20     *SourceShopifyUpdateShopifyAuthorizationMethodOAuth20

	Type SourceShopifyUpdateShopifyAuthorizationMethodType
}

func CreateSourceShopifyUpdateShopifyAuthorizationMethodSourceShopifyUpdateShopifyAuthorizationMethodAPIPassword(sourceShopifyUpdateShopifyAuthorizationMethodAPIPassword SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword) SourceShopifyUpdateShopifyAuthorizationMethod {
	typ := SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodAPIPassword

	return SourceShopifyUpdateShopifyAuthorizationMethod{
		SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword: &sourceShopifyUpdateShopifyAuthorizationMethodAPIPassword,
		Type: typ,
	}
}

func CreateSourceShopifyUpdateShopifyAuthorizationMethodSourceShopifyUpdateShopifyAuthorizationMethodOAuth20(sourceShopifyUpdateShopifyAuthorizationMethodOAuth20 SourceShopifyUpdateShopifyAuthorizationMethodOAuth20) SourceShopifyUpdateShopifyAuthorizationMethod {
	typ := SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodOAuth20

	return SourceShopifyUpdateShopifyAuthorizationMethod{
		SourceShopifyUpdateShopifyAuthorizationMethodOAuth20: &sourceShopifyUpdateShopifyAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func (u *SourceShopifyUpdateShopifyAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceShopifyUpdateShopifyAuthorizationMethodAPIPassword := new(SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceShopifyUpdateShopifyAuthorizationMethodAPIPassword); err == nil {
		u.SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword = sourceShopifyUpdateShopifyAuthorizationMethodAPIPassword
		u.Type = SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodAPIPassword
		return nil
	}

	sourceShopifyUpdateShopifyAuthorizationMethodOAuth20 := new(SourceShopifyUpdateShopifyAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceShopifyUpdateShopifyAuthorizationMethodOAuth20); err == nil {
		u.SourceShopifyUpdateShopifyAuthorizationMethodOAuth20 = sourceShopifyUpdateShopifyAuthorizationMethodOAuth20
		u.Type = SourceShopifyUpdateShopifyAuthorizationMethodTypeSourceShopifyUpdateShopifyAuthorizationMethodOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceShopifyUpdateShopifyAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword != nil {
		return json.Marshal(u.SourceShopifyUpdateShopifyAuthorizationMethodAPIPassword)
	}

	if u.SourceShopifyUpdateShopifyAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceShopifyUpdateShopifyAuthorizationMethodOAuth20)
	}

	return nil, nil
}

type SourceShopifyUpdate struct {
	// The authorization method to use to retrieve data from Shopify
	Credentials *SourceShopifyUpdateShopifyAuthorizationMethod `json:"credentials,omitempty"`
	// The name of your Shopify store found in the URL. For example, if your URL was https://NAME.myshopify.com, then the name would be 'NAME'.
	Shop string `json:"shop"`
	// The date you would like to replicate data from. Format: YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}