// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBingAds struct {
	AuthMethod       types.String `tfsdk:"auth_method"`
	ClientID         types.String `tfsdk:"client_id"`
	ClientSecret     types.String `tfsdk:"client_secret"`
	DeveloperToken   types.String `tfsdk:"developer_token"`
	RefreshToken     types.String `tfsdk:"refresh_token"`
	ReportsStartDate types.String `tfsdk:"reports_start_date"`
	SourceType       types.String `tfsdk:"source_type"`
	TenantID         types.String `tfsdk:"tenant_id"`
}