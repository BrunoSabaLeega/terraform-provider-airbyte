// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsUpdateOutputFormatAvroApacheAvro struct {
	CompressionCodec DestinationGcsUpdateOutputFormatAvroApacheAvroCompressionCodec `tfsdk:"compression_codec"`
	FormatType       types.String                                                   `tfsdk:"format_type"`
}