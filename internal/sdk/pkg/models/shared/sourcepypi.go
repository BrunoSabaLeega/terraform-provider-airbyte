// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePypiPypiEnum string

const (
	SourcePypiPypiEnumPypi SourcePypiPypiEnum = "pypi"
)

func (e *SourcePypiPypiEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pypi":
		*e = SourcePypiPypiEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePypiPypiEnum: %s", s)
	}
}

// SourcePypi - The values required to configure the source.
type SourcePypi struct {
	// Name of the project/package. Can only be in lowercase with hyphen. This is the name used using pip command for installing the package.
	ProjectName string             `json:"project_name"`
	SourceType  SourcePypiPypiEnum `json:"sourceType"`
	// Version of the project/package.  Use it to find a particular release instead of all releases.
	Version *string `json:"version,omitempty"`
}