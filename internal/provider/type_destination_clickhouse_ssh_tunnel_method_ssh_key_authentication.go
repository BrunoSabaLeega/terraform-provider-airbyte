// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication struct {
	SSHKey       types.String `tfsdk:"ssh_key"`
	TunnelHost   types.String `tfsdk:"tunnel_host"`
	TunnelMethod types.String `tfsdk:"tunnel_method"`
	TunnelPort   types.Int64  `tfsdk:"tunnel_port"`
	TunnelUser   types.String `tfsdk:"tunnel_user"`
}
