// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SourceDynamodbDynamodbRegionEnum - The region of the Dynamodb database
type SourceDynamodbDynamodbRegionEnum string

const (
	SourceDynamodbDynamodbRegionEnumUnknown      SourceDynamodbDynamodbRegionEnum = ""
	SourceDynamodbDynamodbRegionEnumUsEast1      SourceDynamodbDynamodbRegionEnum = "us-east-1"
	SourceDynamodbDynamodbRegionEnumUsEast2      SourceDynamodbDynamodbRegionEnum = "us-east-2"
	SourceDynamodbDynamodbRegionEnumUsWest1      SourceDynamodbDynamodbRegionEnum = "us-west-1"
	SourceDynamodbDynamodbRegionEnumUsWest2      SourceDynamodbDynamodbRegionEnum = "us-west-2"
	SourceDynamodbDynamodbRegionEnumAfSouth1     SourceDynamodbDynamodbRegionEnum = "af-south-1"
	SourceDynamodbDynamodbRegionEnumApEast1      SourceDynamodbDynamodbRegionEnum = "ap-east-1"
	SourceDynamodbDynamodbRegionEnumApSouth1     SourceDynamodbDynamodbRegionEnum = "ap-south-1"
	SourceDynamodbDynamodbRegionEnumApNortheast1 SourceDynamodbDynamodbRegionEnum = "ap-northeast-1"
	SourceDynamodbDynamodbRegionEnumApNortheast2 SourceDynamodbDynamodbRegionEnum = "ap-northeast-2"
	SourceDynamodbDynamodbRegionEnumApNortheast3 SourceDynamodbDynamodbRegionEnum = "ap-northeast-3"
	SourceDynamodbDynamodbRegionEnumApSoutheast1 SourceDynamodbDynamodbRegionEnum = "ap-southeast-1"
	SourceDynamodbDynamodbRegionEnumApSoutheast2 SourceDynamodbDynamodbRegionEnum = "ap-southeast-2"
	SourceDynamodbDynamodbRegionEnumCaCentral1   SourceDynamodbDynamodbRegionEnum = "ca-central-1"
	SourceDynamodbDynamodbRegionEnumCnNorth1     SourceDynamodbDynamodbRegionEnum = "cn-north-1"
	SourceDynamodbDynamodbRegionEnumCnNorthwest1 SourceDynamodbDynamodbRegionEnum = "cn-northwest-1"
	SourceDynamodbDynamodbRegionEnumEuCentral1   SourceDynamodbDynamodbRegionEnum = "eu-central-1"
	SourceDynamodbDynamodbRegionEnumEuNorth1     SourceDynamodbDynamodbRegionEnum = "eu-north-1"
	SourceDynamodbDynamodbRegionEnumEuSouth1     SourceDynamodbDynamodbRegionEnum = "eu-south-1"
	SourceDynamodbDynamodbRegionEnumEuWest1      SourceDynamodbDynamodbRegionEnum = "eu-west-1"
	SourceDynamodbDynamodbRegionEnumEuWest2      SourceDynamodbDynamodbRegionEnum = "eu-west-2"
	SourceDynamodbDynamodbRegionEnumEuWest3      SourceDynamodbDynamodbRegionEnum = "eu-west-3"
	SourceDynamodbDynamodbRegionEnumSaEast1      SourceDynamodbDynamodbRegionEnum = "sa-east-1"
	SourceDynamodbDynamodbRegionEnumMeSouth1     SourceDynamodbDynamodbRegionEnum = "me-south-1"
	SourceDynamodbDynamodbRegionEnumUsGovEast1   SourceDynamodbDynamodbRegionEnum = "us-gov-east-1"
	SourceDynamodbDynamodbRegionEnumUsGovWest1   SourceDynamodbDynamodbRegionEnum = "us-gov-west-1"
)

func (e *SourceDynamodbDynamodbRegionEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = SourceDynamodbDynamodbRegionEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDynamodbDynamodbRegionEnum: %s", s)
	}
}

type SourceDynamodbDynamodbEnum string

const (
	SourceDynamodbDynamodbEnumDynamodb SourceDynamodbDynamodbEnum = "dynamodb"
)

func (e *SourceDynamodbDynamodbEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "dynamodb":
		*e = SourceDynamodbDynamodbEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDynamodbDynamodbEnum: %s", s)
	}
}

// SourceDynamodb - The values required to configure the source.
type SourceDynamodb struct {
	// The access key id to access Dynamodb. Airbyte requires read permissions to the database
	AccessKeyID string `json:"access_key_id"`
	// the URL of the Dynamodb database
	Endpoint *string `json:"endpoint,omitempty"`
	// The region of the Dynamodb database
	Region *SourceDynamodbDynamodbRegionEnum `json:"region,omitempty"`
	// Comma separated reserved attribute names present in your tables
	ReservedAttributeNames *string `json:"reserved_attribute_names,omitempty"`
	// The corresponding secret to the access key id.
	SecretAccessKey string                     `json:"secret_access_key"`
	SourceType      SourceDynamodbDynamodbEnum `json:"sourceType"`
}