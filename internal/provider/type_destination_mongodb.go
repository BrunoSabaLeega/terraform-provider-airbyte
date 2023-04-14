// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMongodb struct {
	AuthType        DestinationMongodbAuthorizationType    `tfsdk:"auth_type"`
	Database        types.String                           `tfsdk:"database"`
	DestinationType types.String                           `tfsdk:"destination_type"`
	InstanceType    *DestinationMongodbMongoDbInstanceType `tfsdk:"instance_type"`
	TunnelMethod    *DestinationMongodbSSHTunnelMethod     `tfsdk:"tunnel_method"`
}