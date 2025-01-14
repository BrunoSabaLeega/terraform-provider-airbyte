// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceCoinAPIEnvironment - The environment to use. Either sandbox or production.
type SourceCoinAPIEnvironment string

const (
	SourceCoinAPIEnvironmentSandbox    SourceCoinAPIEnvironment = "sandbox"
	SourceCoinAPIEnvironmentProduction SourceCoinAPIEnvironment = "production"
)

func (e SourceCoinAPIEnvironment) ToPointer() *SourceCoinAPIEnvironment {
	return &e
}

func (e *SourceCoinAPIEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sandbox":
		fallthrough
	case "production":
		*e = SourceCoinAPIEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCoinAPIEnvironment: %v", v)
	}
}

type SourceCoinAPICoinAPI string

const (
	SourceCoinAPICoinAPICoinAPI SourceCoinAPICoinAPI = "coin-api"
)

func (e SourceCoinAPICoinAPI) ToPointer() *SourceCoinAPICoinAPI {
	return &e
}

func (e *SourceCoinAPICoinAPI) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "coin-api":
		*e = SourceCoinAPICoinAPI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCoinAPICoinAPI: %v", v)
	}
}

type SourceCoinAPI struct {
	// API Key
	APIKey string `json:"api_key"`
	// The end date in ISO 8601 format. If not supplied, data will be returned
	// from the start date to the current time, or when the count of result
	// elements reaches its limit.
	//
	EndDate *string `json:"end_date,omitempty"`
	// The environment to use. Either sandbox or production.
	//
	Environment SourceCoinAPIEnvironment `json:"environment"`
	// The maximum number of elements to return. If not supplied, the default
	// is 100. For numbers larger than 100, each 100 items is counted as one
	// request for pricing purposes. Maximum value is 100000.
	//
	Limit *int64 `json:"limit,omitempty"`
	// The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get
	Period     string               `json:"period"`
	SourceType SourceCoinAPICoinAPI `json:"sourceType"`
	// The start date in ISO 8601 format.
	StartDate string `json:"start_date"`
	// The symbol ID to use. See the documentation for a list.
	// https://docs.coinapi.io/#list-all-symbols-get
	//
	SymbolID string `json:"symbol_id"`
}
