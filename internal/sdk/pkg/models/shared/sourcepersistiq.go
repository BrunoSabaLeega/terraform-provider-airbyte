// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePersistiqPersistiqEnum string

const (
	SourcePersistiqPersistiqEnumPersistiq SourcePersistiqPersistiqEnum = "persistiq"
)

func (e *SourcePersistiqPersistiqEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "persistiq":
		*e = SourcePersistiqPersistiqEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePersistiqPersistiqEnum: %s", s)
	}
}

// SourcePersistiq - The values required to configure the source.
type SourcePersistiq struct {
	// PersistIq API Key. See the <a href="https://apidocs.persistiq.com/#authentication">docs</a> for more information on where to find that key.
	APIKey     string                       `json:"api_key"`
	SourceType SourcePersistiqPersistiqEnum `json:"sourceType"`
}