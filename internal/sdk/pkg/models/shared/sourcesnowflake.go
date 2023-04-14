// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum string

const (
	SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnumUsernamePassword SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum = "username/password"
)

func (e *SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "username/password":
		*e = SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum: %s", s)
	}
}

type SourceSnowflakeAuthorizationMethodUsernameAndPassword struct {
	AuthType SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum `json:"auth_type"`
	// The password associated with the username.
	Password string `json:"password"`
	// The username you created to allow Airbyte to access the database.
	Username string `json:"username"`
}

type SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum string

const (
	SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnumOAuth SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum = "OAuth"
)

func (e *SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "OAuth":
		*e = SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum: %s", s)
	}
}

type SourceSnowflakeAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string                                               `json:"access_token,omitempty"`
	AuthType    SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum `json:"auth_type"`
	// The Client ID of your Snowflake developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Snowflake developer application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token for making authenticated requests.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type SourceSnowflakeAuthorizationMethodType string

const (
	SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodOAuth20             SourceSnowflakeAuthorizationMethodType = "source-snowflake_Authorization Method_OAuth2.0"
	SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodUsernameAndPassword SourceSnowflakeAuthorizationMethodType = "source-snowflake_Authorization Method_Username and Password"
)

type SourceSnowflakeAuthorizationMethod struct {
	SourceSnowflakeAuthorizationMethodOAuth20             *SourceSnowflakeAuthorizationMethodOAuth20
	SourceSnowflakeAuthorizationMethodUsernameAndPassword *SourceSnowflakeAuthorizationMethodUsernameAndPassword

	Type SourceSnowflakeAuthorizationMethodType
}

func CreateSourceSnowflakeAuthorizationMethodSourceSnowflakeAuthorizationMethodOAuth20(sourceSnowflakeAuthorizationMethodOAuth20 SourceSnowflakeAuthorizationMethodOAuth20) SourceSnowflakeAuthorizationMethod {
	typ := SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodOAuth20

	return SourceSnowflakeAuthorizationMethod{
		SourceSnowflakeAuthorizationMethodOAuth20: &sourceSnowflakeAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceSnowflakeAuthorizationMethodSourceSnowflakeAuthorizationMethodUsernameAndPassword(sourceSnowflakeAuthorizationMethodUsernameAndPassword SourceSnowflakeAuthorizationMethodUsernameAndPassword) SourceSnowflakeAuthorizationMethod {
	typ := SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodUsernameAndPassword

	return SourceSnowflakeAuthorizationMethod{
		SourceSnowflakeAuthorizationMethodUsernameAndPassword: &sourceSnowflakeAuthorizationMethodUsernameAndPassword,
		Type: typ,
	}
}

func (u *SourceSnowflakeAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSnowflakeAuthorizationMethodOAuth20 := new(SourceSnowflakeAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSnowflakeAuthorizationMethodOAuth20); err == nil {
		u.SourceSnowflakeAuthorizationMethodOAuth20 = sourceSnowflakeAuthorizationMethodOAuth20
		u.Type = SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodOAuth20
		return nil
	}

	sourceSnowflakeAuthorizationMethodUsernameAndPassword := new(SourceSnowflakeAuthorizationMethodUsernameAndPassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSnowflakeAuthorizationMethodUsernameAndPassword); err == nil {
		u.SourceSnowflakeAuthorizationMethodUsernameAndPassword = sourceSnowflakeAuthorizationMethodUsernameAndPassword
		u.Type = SourceSnowflakeAuthorizationMethodTypeSourceSnowflakeAuthorizationMethodUsernameAndPassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSnowflakeAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceSnowflakeAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceSnowflakeAuthorizationMethodOAuth20)
	}

	if u.SourceSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		return json.Marshal(u.SourceSnowflakeAuthorizationMethodUsernameAndPassword)
	}

	return nil, nil
}

type SourceSnowflakeSnowflakeEnum string

const (
	SourceSnowflakeSnowflakeEnumSnowflake SourceSnowflakeSnowflakeEnum = "snowflake"
)

func (e *SourceSnowflakeSnowflakeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "snowflake":
		*e = SourceSnowflakeSnowflakeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeSnowflakeEnum: %s", s)
	}
}

// SourceSnowflake - The values required to configure the source.
type SourceSnowflake struct {
	Credentials *SourceSnowflakeAuthorizationMethod `json:"credentials,omitempty"`
	// The database you created for Airbyte to access data.
	Database string `json:"database"`
	// The host domain of the snowflake instance (must include the account, region, cloud environment, and end with snowflakecomputing.com).
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The role you created for Airbyte to access Snowflake.
	Role string `json:"role"`
	// The source Snowflake schema tables. Leave empty to access tables from multiple schemas.
	Schema     *string                      `json:"schema,omitempty"`
	SourceType SourceSnowflakeSnowflakeEnum `json:"sourceType"`
	// The warehouse you created for Airbyte to access data.
	Warehouse string `json:"warehouse"`
}