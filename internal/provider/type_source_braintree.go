// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBraintree struct {
	Environment types.String `tfsdk:"environment"`
	MerchantID  types.String `tfsdk:"merchant_id"`
	PrivateKey  types.String `tfsdk:"private_key"`
	PublicKey   types.String `tfsdk:"public_key"`
	SourceType  types.String `tfsdk:"source_type"`
	StartDate   types.String `tfsdk:"start_date"`
}
