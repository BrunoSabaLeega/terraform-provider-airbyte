// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationLangchainIndexingChromaLocalPersistance struct {
	CollectionName  types.String `tfsdk:"collection_name"`
	DestinationPath types.String `tfsdk:"destination_path"`
	Mode            types.String `tfsdk:"mode"`
}
