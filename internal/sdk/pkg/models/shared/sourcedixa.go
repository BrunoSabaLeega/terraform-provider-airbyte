// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceDixaDixaEnum string

const (
	SourceDixaDixaEnumDixa SourceDixaDixaEnum = "dixa"
)

func (e *SourceDixaDixaEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "dixa":
		*e = SourceDixaDixaEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDixaDixaEnum: %s", s)
	}
}

// SourceDixa - The values required to configure the source.
type SourceDixa struct {
	// Dixa API token
	APIToken string `json:"api_token"`
	// Number of days to batch into one request. Max 31.
	BatchSize  *int64             `json:"batch_size,omitempty"`
	SourceType SourceDixaDixaEnum `json:"sourceType"`
	// The connector pulls records updated from this date onwards.
	StartDate string `json:"start_date"`
}