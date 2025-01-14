// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMilvusIndexing struct {
	Auth        DestinationMilvusIndexingAuthentication `tfsdk:"auth"`
	Collection  types.String                            `tfsdk:"collection"`
	Db          types.String                            `tfsdk:"db"`
	Host        types.String                            `tfsdk:"host"`
	TextField   types.String                            `tfsdk:"text_field"`
	VectorField types.String                            `tfsdk:"vector_field"`
}
