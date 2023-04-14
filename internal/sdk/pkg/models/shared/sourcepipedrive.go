// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourcePipedriveAPIKeyAuthenticationAuthTypeEnum string

const (
	SourcePipedriveAPIKeyAuthenticationAuthTypeEnumToken SourcePipedriveAPIKeyAuthenticationAuthTypeEnum = "Token"
)

func (e *SourcePipedriveAPIKeyAuthenticationAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Token":
		*e = SourcePipedriveAPIKeyAuthenticationAuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePipedriveAPIKeyAuthenticationAuthTypeEnum: %s", s)
	}
}

type SourcePipedriveAPIKeyAuthentication struct {
	// The Pipedrive API Token.
	APIToken string                                          `json:"api_token"`
	AuthType SourcePipedriveAPIKeyAuthenticationAuthTypeEnum `json:"auth_type"`
}

type SourcePipedrivePipedriveEnum string

const (
	SourcePipedrivePipedriveEnumPipedrive SourcePipedrivePipedriveEnum = "pipedrive"
)

func (e *SourcePipedrivePipedriveEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pipedrive":
		*e = SourcePipedrivePipedriveEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePipedrivePipedriveEnum: %s", s)
	}
}

// SourcePipedrive - The values required to configure the source.
type SourcePipedrive struct {
	Authorization *SourcePipedriveAPIKeyAuthentication `json:"authorization,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. When specified and not None, then stream will behave as incremental
	ReplicationStartDate time.Time                    `json:"replication_start_date"`
	SourceType           SourcePipedrivePipedriveEnum `json:"sourceType"`
}