// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceBigcommerceBigcommerce string

const (
	SourceBigcommerceBigcommerceBigcommerce SourceBigcommerceBigcommerce = "bigcommerce"
)

func (e SourceBigcommerceBigcommerce) ToPointer() *SourceBigcommerceBigcommerce {
	return &e
}

func (e *SourceBigcommerceBigcommerce) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bigcommerce":
		*e = SourceBigcommerceBigcommerce(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBigcommerceBigcommerce: %v", v)
	}
}

type SourceBigcommerce struct {
	// Access Token for making authenticated requests.
	AccessToken string                       `json:"access_token"`
	SourceType  SourceBigcommerceBigcommerce `json:"sourceType"`
	// The date you would like to replicate data. Format: YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The hash code of the store. For https://api.bigcommerce.com/stores/HASH_CODE/v3/, The store's hash code is 'HASH_CODE'.
	StoreHash string `json:"store_hash"`
}
