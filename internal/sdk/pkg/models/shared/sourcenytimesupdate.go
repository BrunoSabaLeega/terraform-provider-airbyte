// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

// SourceNytimesUpdatePeriodUsedForMostPopularStreams - Period of time (in days)
type SourceNytimesUpdatePeriodUsedForMostPopularStreams int64

const (
	SourceNytimesUpdatePeriodUsedForMostPopularStreamsOne    SourceNytimesUpdatePeriodUsedForMostPopularStreams = 1
	SourceNytimesUpdatePeriodUsedForMostPopularStreamsSeven  SourceNytimesUpdatePeriodUsedForMostPopularStreams = 7
	SourceNytimesUpdatePeriodUsedForMostPopularStreamsThirty SourceNytimesUpdatePeriodUsedForMostPopularStreams = 30
)

func (e SourceNytimesUpdatePeriodUsedForMostPopularStreams) ToPointer() *SourceNytimesUpdatePeriodUsedForMostPopularStreams {
	return &e
}

func (e *SourceNytimesUpdatePeriodUsedForMostPopularStreams) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 30:
		*e = SourceNytimesUpdatePeriodUsedForMostPopularStreams(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNytimesUpdatePeriodUsedForMostPopularStreams: %v", v)
	}
}

// SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream - Share Type
type SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream string

const (
	SourceNytimesUpdateShareTypeUsedForMostPopularSharedStreamFacebook SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream = "facebook"
)

func (e SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream) ToPointer() *SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream {
	return &e
}

func (e *SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "facebook":
		*e = SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream: %v", v)
	}
}

type SourceNytimesUpdate struct {
	// API Key
	APIKey string `json:"api_key"`
	// End date to stop the article retrieval (format YYYY-MM)
	EndDate *types.Date `json:"end_date,omitempty"`
	// Period of time (in days)
	Period SourceNytimesUpdatePeriodUsedForMostPopularStreams `json:"period"`
	// Share Type
	ShareType *SourceNytimesUpdateShareTypeUsedForMostPopularSharedStream `json:"share_type,omitempty"`
	// Start date to begin the article retrieval (format YYYY-MM)
	StartDate types.Date `json:"start_date"`
}