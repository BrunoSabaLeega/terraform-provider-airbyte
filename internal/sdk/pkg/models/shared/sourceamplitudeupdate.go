// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceAmplitudeUpdateDataRegion - Amplitude data region server
type SourceAmplitudeUpdateDataRegion string

const (
	SourceAmplitudeUpdateDataRegionStandardServer    SourceAmplitudeUpdateDataRegion = "Standard Server"
	SourceAmplitudeUpdateDataRegionEuResidencyServer SourceAmplitudeUpdateDataRegion = "EU Residency Server"
)

func (e SourceAmplitudeUpdateDataRegion) ToPointer() *SourceAmplitudeUpdateDataRegion {
	return &e
}

func (e *SourceAmplitudeUpdateDataRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard Server":
		fallthrough
	case "EU Residency Server":
		*e = SourceAmplitudeUpdateDataRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmplitudeUpdateDataRegion: %v", v)
	}
}

type SourceAmplitudeUpdate struct {
	// Amplitude API Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Amplitude data region server
	DataRegion *SourceAmplitudeUpdateDataRegion `json:"data_region,omitempty"`
	// According to <a href="https://www.docs.developers.amplitude.com/analytics/apis/export-api/#considerations">Considerations</a> too big time range in request can cause a timeout error. In this case, set shorter time interval in hours.
	RequestTimeRange *int64 `json:"request_time_range,omitempty"`
	// Amplitude Secret Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
	SecretKey string `json:"secret_key"`
	// UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}
