// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationKeen struct {
	APIKey          types.String `tfsdk:"api_key"`
	DestinationType types.String `tfsdk:"destination_type"`
	InferTimestamp  types.Bool   `tfsdk:"infer_timestamp"`
	ProjectID       types.String `tfsdk:"project_id"`
}
