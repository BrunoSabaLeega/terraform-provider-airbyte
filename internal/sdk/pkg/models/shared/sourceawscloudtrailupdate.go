// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
)

type SourceAwsCloudtrailUpdate struct {
	// AWS CloudTrail Access Key ID. See the <a href="https://docs.airbyte.com/integrations/sources/aws-cloudtrail">docs</a> for more information on how to obtain this key.
	AwsKeyID string `json:"aws_key_id"`
	// The default AWS Region to use, for example, us-west-1 or us-west-2. When specifying a Region inline during client initialization, this property is named region_name.
	AwsRegionName string `json:"aws_region_name"`
	// AWS CloudTrail Access Key ID. See the <a href="https://docs.airbyte.com/integrations/sources/aws-cloudtrail">docs</a> for more information on how to obtain this key.
	AwsSecretKey string `json:"aws_secret_key"`
	// The date you would like to replicate data. Data in AWS CloudTrail is available for last 90 days only. Format: YYYY-MM-DD.
	StartDate types.Date `json:"start_date"`
}