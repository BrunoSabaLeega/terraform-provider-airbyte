// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAwsCloudtrail struct {
	AwsKeyID      types.String `tfsdk:"aws_key_id"`
	AwsRegionName types.String `tfsdk:"aws_region_name"`
	AwsSecretKey  types.String `tfsdk:"aws_secret_key"`
	SourceType    types.String `tfsdk:"source_type"`
	StartDate     types.String `tfsdk:"start_date"`
}
