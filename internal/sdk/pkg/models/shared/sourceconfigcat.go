// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceConfigcatConfigcat string

const (
	SourceConfigcatConfigcatConfigcat SourceConfigcatConfigcat = "configcat"
)

func (e SourceConfigcatConfigcat) ToPointer() *SourceConfigcatConfigcat {
	return &e
}

func (e *SourceConfigcatConfigcat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "configcat":
		*e = SourceConfigcatConfigcat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceConfigcatConfigcat: %v", v)
	}
}

type SourceConfigcat struct {
	// Basic auth password. See <a href="https://api.configcat.com/docs/#section/Authentication">here</a>.
	Password   string                   `json:"password"`
	SourceType SourceConfigcatConfigcat `json:"sourceType"`
	// Basic auth user name. See <a href="https://api.configcat.com/docs/#section/Authentication">here</a>.
	Username string `json:"username"`
}
