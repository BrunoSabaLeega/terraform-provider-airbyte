// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationKinesisKinesisEnum string

const (
	DestinationKinesisKinesisEnumKinesis DestinationKinesisKinesisEnum = "kinesis"
)

func (e *DestinationKinesisKinesisEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "kinesis":
		*e = DestinationKinesisKinesisEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationKinesisKinesisEnum: %s", s)
	}
}

// DestinationKinesis - The values required to configure the destination.
type DestinationKinesis struct {
	// Generate the AWS Access Key for current user.
	AccessKey string `json:"accessKey"`
	// Buffer size for storing kinesis records before being batch streamed.
	BufferSize      int64                         `json:"bufferSize"`
	DestinationType DestinationKinesisKinesisEnum `json:"destinationType"`
	// AWS Kinesis endpoint.
	Endpoint string `json:"endpoint"`
	// The AWS Private Key - a string of numbers and letters that are unique for each account, also known as a "recovery phrase".
	PrivateKey string `json:"privateKey"`
	// AWS region. Your account determines the Regions that are available to you.
	Region string `json:"region"`
	// Number of shards to which the data should be streamed.
	ShardCount int64 `json:"shardCount"`
}