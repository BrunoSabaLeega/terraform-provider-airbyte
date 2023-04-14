// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceStrava struct {
	AthleteID    types.Int64  `tfsdk:"athlete_id"`
	AuthType     types.String `tfsdk:"auth_type"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	RefreshToken types.String `tfsdk:"refresh_token"`
	SourceType   types.String `tfsdk:"source_type"`
	StartDate    types.String `tfsdk:"start_date"`
}