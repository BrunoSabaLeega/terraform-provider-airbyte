//source_kafka_resource

package provider

import (
	"airbyte/internal/sdk"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceKafkaResource{}
var _ resource.ResourceWithImportState = &SourceKafkaResource{}

func NewSourceKafkaResource() resource.Resource {
	return &SourceKafkaResource{}
}

// SourceKafkaResource defines the resource implementation.
type SourceKafkaResource struct {
	client *sdk.SDK
}

// SourceKafkaResourceModel describes the resource data model.
//type SourceKafkaResourceModel struct {
//	Configuration SourceKafka `tfsdk:"configuration"`
//	Name          types.String `tfsdk:"name"`
//	SecretID      types.String `tfsdk:"secret_id"`
//	SourceID      types.String `tfsdk:"source_id"`
//	SourceType    types.String `tfsdk:"source_type"`
//	WorkspaceID   types.String `tfsdk:"workspace_id"`
//}

func (r *SourceKafkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_kafka"
}