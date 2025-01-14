// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGooglePagespeedInsights struct {
	APIKey     types.String   `tfsdk:"api_key"`
	Categories []types.String `tfsdk:"categories"`
	SourceType types.String   `tfsdk:"source_type"`
	Strategies []types.String `tfsdk:"strategies"`
	Urls       []types.String `tfsdk:"urls"`
}
