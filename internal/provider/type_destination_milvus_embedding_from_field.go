// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMilvusEmbeddingFromField struct {
	Dimensions types.Int64  `tfsdk:"dimensions"`
	FieldName  types.String `tfsdk:"field_name"`
	Mode       types.String `tfsdk:"mode"`
}
