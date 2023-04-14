// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationKinesis struct {
	AccessKey       types.String `tfsdk:"access_key"`
	BufferSize      types.Int64  `tfsdk:"buffer_size"`
	DestinationType types.String `tfsdk:"destination_type"`
	Endpoint        types.String `tfsdk:"endpoint"`
	PrivateKey      types.String `tfsdk:"private_key"`
	Region          types.String `tfsdk:"region"`
	ShardCount      types.Int64  `tfsdk:"shard_count"`
}