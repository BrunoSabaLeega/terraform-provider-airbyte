// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceLaunchdarklyLaunchdarklyEnum string

const (
	SourceLaunchdarklyLaunchdarklyEnumLaunchdarkly SourceLaunchdarklyLaunchdarklyEnum = "launchdarkly"
)

func (e *SourceLaunchdarklyLaunchdarklyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "launchdarkly":
		*e = SourceLaunchdarklyLaunchdarklyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLaunchdarklyLaunchdarklyEnum: %s", s)
	}
}

// SourceLaunchdarkly - The values required to configure the source.
type SourceLaunchdarkly struct {
	// Your Access token. See <a href="https://apidocs.launchdarkly.com/#section/Overview/Authentication">here</a>.
	AccessToken string                             `json:"access_token"`
	SourceType  SourceLaunchdarklyLaunchdarklyEnum `json:"sourceType"`
}