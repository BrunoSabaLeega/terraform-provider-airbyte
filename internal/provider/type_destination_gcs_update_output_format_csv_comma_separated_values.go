// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsUpdateOutputFormatCSVCommaSeparatedValues struct {
	Compression *DestinationGcsUpdateOutputFormatCSVCommaSeparatedValuesCompression `tfsdk:"compression"`
	Flattening  types.String                                                        `tfsdk:"flattening"`
	FormatType  types.String                                                        `tfsdk:"format_type"`
}
