// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationLangchainIndexingPinecone struct {
	Index               types.String `tfsdk:"index"`
	Mode                types.String `tfsdk:"mode"`
	PineconeEnvironment types.String `tfsdk:"pinecone_environment"`
	PineconeKey         types.String `tfsdk:"pinecone_key"`
}