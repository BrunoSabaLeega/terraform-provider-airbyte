// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum - Name of the credentials
type DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum string

const (
	DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnumIamUser DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum = "IAM User"
)

func (e *DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "IAM User":
		*e = DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum: %s", s)
	}
}

// DestinationAwsDatalakeAuthenticationModeIAMUser - Choose How to Authenticate to AWS.
type DestinationAwsDatalakeAuthenticationModeIAMUser struct {
	// AWS User Access Key Id
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// Secret Access Key
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// Name of the credentials
	CredentialsTitle DestinationAwsDatalakeAuthenticationModeIAMUserCredentialsTitleEnum `json:"credentials_title"`
}

// DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum - Name of the credentials
type DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum string

const (
	DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnumIamRole DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum = "IAM Role"
)

func (e *DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "IAM Role":
		*e = DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum: %s", s)
	}
}

// DestinationAwsDatalakeAuthenticationModeIAMRole - Choose How to Authenticate to AWS.
type DestinationAwsDatalakeAuthenticationModeIAMRole struct {
	// Name of the credentials
	CredentialsTitle DestinationAwsDatalakeAuthenticationModeIAMRoleCredentialsTitleEnum `json:"credentials_title"`
	// Will assume this role to write data to s3
	RoleArn string `json:"role_arn"`
}

type DestinationAwsDatalakeAuthenticationModeType string

const (
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_Authentication mode_IAM Role"
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_Authentication mode_IAM User"
)

type DestinationAwsDatalakeAuthenticationMode struct {
	DestinationAwsDatalakeAuthenticationModeIAMRole *DestinationAwsDatalakeAuthenticationModeIAMRole
	DestinationAwsDatalakeAuthenticationModeIAMUser *DestinationAwsDatalakeAuthenticationModeIAMUser

	Type DestinationAwsDatalakeAuthenticationModeType
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeAuthenticationModeIAMRole(destinationAwsDatalakeAuthenticationModeIAMRole DestinationAwsDatalakeAuthenticationModeIAMRole) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeAuthenticationModeIAMRole: &destinationAwsDatalakeAuthenticationModeIAMRole,
		Type: typ,
	}
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeAuthenticationModeIAMUser(destinationAwsDatalakeAuthenticationModeIAMUser DestinationAwsDatalakeAuthenticationModeIAMUser) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeAuthenticationModeIAMUser: &destinationAwsDatalakeAuthenticationModeIAMUser,
		Type: typ,
	}
}

func (u *DestinationAwsDatalakeAuthenticationMode) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationAwsDatalakeAuthenticationModeIAMRole := new(DestinationAwsDatalakeAuthenticationModeIAMRole)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeAuthenticationModeIAMRole); err == nil {
		u.DestinationAwsDatalakeAuthenticationModeIAMRole = destinationAwsDatalakeAuthenticationModeIAMRole
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMRole
		return nil
	}

	destinationAwsDatalakeAuthenticationModeIAMUser := new(DestinationAwsDatalakeAuthenticationModeIAMUser)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeAuthenticationModeIAMUser); err == nil {
		u.DestinationAwsDatalakeAuthenticationModeIAMUser = destinationAwsDatalakeAuthenticationModeIAMUser
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeAuthenticationModeIAMUser
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeAuthenticationMode) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeAuthenticationModeIAMRole != nil {
		return json.Marshal(u.DestinationAwsDatalakeAuthenticationModeIAMRole)
	}

	if u.DestinationAwsDatalakeAuthenticationModeIAMUser != nil {
		return json.Marshal(u.DestinationAwsDatalakeAuthenticationModeIAMUser)
	}

	return nil, nil
}

type DestinationAwsDatalakeAwsDatalakeEnum string

const (
	DestinationAwsDatalakeAwsDatalakeEnumAwsDatalake DestinationAwsDatalakeAwsDatalakeEnum = "aws-datalake"
)

func (e *DestinationAwsDatalakeAwsDatalakeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "aws-datalake":
		*e = DestinationAwsDatalakeAwsDatalakeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeAwsDatalakeEnum: %s", s)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum - The compression algorithm used to compress data.
type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum string

const (
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnumUncompressed DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum = "UNCOMPRESSED"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnumSnappy       DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum = "SNAPPY"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnumGzip         DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum = "GZIP"
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnumZstd         DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum = "ZSTD"
)

func (e *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "UNCOMPRESSED":
		fallthrough
	case "SNAPPY":
		fallthrough
	case "GZIP":
		fallthrough
	case "ZSTD":
		*e = DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum: %s", s)
	}
}

type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum string

const (
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnumParquet DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum = "Parquet"
)

func (e *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Parquet":
		*e = DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum: %s", s)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage - Format of the data output.
type DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageCompressionCodecOptionalEnum `json:"compression_codec,omitempty"`
	FormatType       DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorageFormatTypeWildcardEnum        `json:"format_type"`
}

// DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum - The compression algorithm used to compress data.
type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum string

const (
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnumUncompressed DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum = "UNCOMPRESSED"
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnumGzip         DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum = "GZIP"
)

func (e *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "UNCOMPRESSED":
		fallthrough
	case "GZIP":
		*e = DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum: %s", s)
	}
}

type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum string

const (
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnumJsonl DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum = "JSONL"
)

func (e *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "JSONL":
		*e = DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum: %s", s)
	}
}

// DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON - Format of the data output.
type DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONCompressionCodecOptionalEnum `json:"compression_codec,omitempty"`
	FormatType       DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSONFormatTypeWildcardEnum        `json:"format_type"`
}

type DestinationAwsDatalakeOutputFormatWildcardType string

const (
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_Output Format *_JSON Lines: Newline-delimited JSON"
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage        DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_Output Format *_Parquet: Columnar Storage"
)

type DestinationAwsDatalakeOutputFormatWildcard struct {
	DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON *DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
	DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage        *DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage

	Type DestinationAwsDatalakeOutputFormatWildcardType
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON(destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON: &destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage(destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage: &destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage,
		Type: typ,
	}
}

func (u *DestinationAwsDatalakeOutputFormatWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON := new(DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON); err == nil {
		u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON = destinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON
		return nil
	}

	destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage := new(DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage); err == nil {
		u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage = destinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeOutputFormatWildcard) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON != nil {
		return json.Marshal(u.DestinationAwsDatalakeOutputFormatWildcardJSONLinesNewlineDelimitedJSON)
	}

	if u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage != nil {
		return json.Marshal(u.DestinationAwsDatalakeOutputFormatWildcardParquetColumnarStorage)
	}

	return nil, nil
}

// DestinationAwsDatalakeChooseHowToPartitionDataEnum - Partition data by cursor fields when a cursor field is a date
type DestinationAwsDatalakeChooseHowToPartitionDataEnum string

const (
	DestinationAwsDatalakeChooseHowToPartitionDataEnumNoPartitioning DestinationAwsDatalakeChooseHowToPartitionDataEnum = "NO PARTITIONING"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumDate           DestinationAwsDatalakeChooseHowToPartitionDataEnum = "DATE"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumYear           DestinationAwsDatalakeChooseHowToPartitionDataEnum = "YEAR"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumMonth          DestinationAwsDatalakeChooseHowToPartitionDataEnum = "MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumDay            DestinationAwsDatalakeChooseHowToPartitionDataEnum = "DAY"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumYearMonth      DestinationAwsDatalakeChooseHowToPartitionDataEnum = "YEAR/MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataEnumYearMonthDay   DestinationAwsDatalakeChooseHowToPartitionDataEnum = "YEAR/MONTH/DAY"
)

func (e *DestinationAwsDatalakeChooseHowToPartitionDataEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "NO PARTITIONING":
		fallthrough
	case "DATE":
		fallthrough
	case "YEAR":
		fallthrough
	case "MONTH":
		fallthrough
	case "DAY":
		fallthrough
	case "YEAR/MONTH":
		fallthrough
	case "YEAR/MONTH/DAY":
		*e = DestinationAwsDatalakeChooseHowToPartitionDataEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeChooseHowToPartitionDataEnum: %s", s)
	}
}

// DestinationAwsDatalakeS3BucketRegionEnum - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationAwsDatalakeS3BucketRegionEnum string

const (
	DestinationAwsDatalakeS3BucketRegionEnumUnknown      DestinationAwsDatalakeS3BucketRegionEnum = ""
	DestinationAwsDatalakeS3BucketRegionEnumUsEast1      DestinationAwsDatalakeS3BucketRegionEnum = "us-east-1"
	DestinationAwsDatalakeS3BucketRegionEnumUsEast2      DestinationAwsDatalakeS3BucketRegionEnum = "us-east-2"
	DestinationAwsDatalakeS3BucketRegionEnumUsWest1      DestinationAwsDatalakeS3BucketRegionEnum = "us-west-1"
	DestinationAwsDatalakeS3BucketRegionEnumUsWest2      DestinationAwsDatalakeS3BucketRegionEnum = "us-west-2"
	DestinationAwsDatalakeS3BucketRegionEnumAfSouth1     DestinationAwsDatalakeS3BucketRegionEnum = "af-south-1"
	DestinationAwsDatalakeS3BucketRegionEnumApEast1      DestinationAwsDatalakeS3BucketRegionEnum = "ap-east-1"
	DestinationAwsDatalakeS3BucketRegionEnumApSouth1     DestinationAwsDatalakeS3BucketRegionEnum = "ap-south-1"
	DestinationAwsDatalakeS3BucketRegionEnumApNortheast1 DestinationAwsDatalakeS3BucketRegionEnum = "ap-northeast-1"
	DestinationAwsDatalakeS3BucketRegionEnumApNortheast2 DestinationAwsDatalakeS3BucketRegionEnum = "ap-northeast-2"
	DestinationAwsDatalakeS3BucketRegionEnumApNortheast3 DestinationAwsDatalakeS3BucketRegionEnum = "ap-northeast-3"
	DestinationAwsDatalakeS3BucketRegionEnumApSoutheast1 DestinationAwsDatalakeS3BucketRegionEnum = "ap-southeast-1"
	DestinationAwsDatalakeS3BucketRegionEnumApSoutheast2 DestinationAwsDatalakeS3BucketRegionEnum = "ap-southeast-2"
	DestinationAwsDatalakeS3BucketRegionEnumCaCentral1   DestinationAwsDatalakeS3BucketRegionEnum = "ca-central-1"
	DestinationAwsDatalakeS3BucketRegionEnumCnNorth1     DestinationAwsDatalakeS3BucketRegionEnum = "cn-north-1"
	DestinationAwsDatalakeS3BucketRegionEnumCnNorthwest1 DestinationAwsDatalakeS3BucketRegionEnum = "cn-northwest-1"
	DestinationAwsDatalakeS3BucketRegionEnumEuCentral1   DestinationAwsDatalakeS3BucketRegionEnum = "eu-central-1"
	DestinationAwsDatalakeS3BucketRegionEnumEuNorth1     DestinationAwsDatalakeS3BucketRegionEnum = "eu-north-1"
	DestinationAwsDatalakeS3BucketRegionEnumEuSouth1     DestinationAwsDatalakeS3BucketRegionEnum = "eu-south-1"
	DestinationAwsDatalakeS3BucketRegionEnumEuWest1      DestinationAwsDatalakeS3BucketRegionEnum = "eu-west-1"
	DestinationAwsDatalakeS3BucketRegionEnumEuWest2      DestinationAwsDatalakeS3BucketRegionEnum = "eu-west-2"
	DestinationAwsDatalakeS3BucketRegionEnumEuWest3      DestinationAwsDatalakeS3BucketRegionEnum = "eu-west-3"
	DestinationAwsDatalakeS3BucketRegionEnumSaEast1      DestinationAwsDatalakeS3BucketRegionEnum = "sa-east-1"
	DestinationAwsDatalakeS3BucketRegionEnumMeSouth1     DestinationAwsDatalakeS3BucketRegionEnum = "me-south-1"
	DestinationAwsDatalakeS3BucketRegionEnumUsGovEast1   DestinationAwsDatalakeS3BucketRegionEnum = "us-gov-east-1"
	DestinationAwsDatalakeS3BucketRegionEnumUsGovWest1   DestinationAwsDatalakeS3BucketRegionEnum = "us-gov-west-1"
)

func (e *DestinationAwsDatalakeS3BucketRegionEnum) UnmarshalJSON(data []byte) error {
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
		*e = DestinationAwsDatalakeS3BucketRegionEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeS3BucketRegionEnum: %s", s)
	}
}

// DestinationAwsDatalake - The values required to configure the destination.
type DestinationAwsDatalake struct {
	// target aws account id
	AwsAccountID *string `json:"aws_account_id,omitempty"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	BucketName string `json:"bucket_name"`
	// S3 prefix
	BucketPrefix *string `json:"bucket_prefix,omitempty"`
	// Choose How to Authenticate to AWS.
	Credentials     DestinationAwsDatalakeAuthenticationMode `json:"credentials"`
	DestinationType DestinationAwsDatalakeAwsDatalakeEnum    `json:"destinationType"`
	// Format of the data output.
	Format *DestinationAwsDatalakeOutputFormatWildcard `json:"format,omitempty"`
	// Cast float/double as decimal(38,18). This can help achieve higher accuracy and represent numbers correctly as received from the source.
	GlueCatalogFloatAsDecimal *bool `json:"glue_catalog_float_as_decimal,omitempty"`
	// Add a default tag key to databases created by this destination
	LakeformationDatabaseDefaultTagKey *string `json:"lakeformation_database_default_tag_key,omitempty"`
	// Add default values for the `Tag Key` to databases created by this destination. Comma separate for multiple values.
	LakeformationDatabaseDefaultTagValues *string `json:"lakeformation_database_default_tag_values,omitempty"`
	// The default database this destination will use to create tables in per stream. Can be changed per connection by customizing the namespace.
	LakeformationDatabaseName string `json:"lakeformation_database_name"`
	// Whether to create tables as LF governed tables.
	LakeformationGovernedTables *bool `json:"lakeformation_governed_tables,omitempty"`
	// Partition data by cursor fields when a cursor field is a date
	Partitioning *DestinationAwsDatalakeChooseHowToPartitionDataEnum `json:"partitioning,omitempty"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	Region DestinationAwsDatalakeS3BucketRegionEnum `json:"region"`
}