// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationBigqueryDenormalized struct {
	BigQueryClientBufferSizeMb types.Int64                                   `tfsdk:"big_query_client_buffer_size_mb"`
	CredentialsJSON            types.String                                  `tfsdk:"credentials_json"`
	DatasetID                  types.String                                  `tfsdk:"dataset_id"`
	DatasetLocation            types.String                                  `tfsdk:"dataset_location"`
	DestinationType            types.String                                  `tfsdk:"destination_type"`
	LoadingMethod              *DestinationBigqueryDenormalizedLoadingMethod `tfsdk:"loading_method"`
	ProjectID                  types.String                                  `tfsdk:"project_id"`
}