// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGcs struct {
	GcsBucket      types.String `tfsdk:"gcs_bucket"`
	GcsPath        types.String `tfsdk:"gcs_path"`
	ServiceAccount types.String `tfsdk:"service_account"`
	SourceType     types.String `tfsdk:"source_type"`
}
