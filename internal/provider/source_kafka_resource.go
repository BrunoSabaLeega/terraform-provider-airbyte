//source_kafka_resource

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/path"
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

func (r *SourceKafkaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceKafkaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	return
}

func (r *SourceKafkaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	return
}

func (r *SourceKafkaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	return
}

func (r *SourceKafkaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	return
}

func (r *SourceKafkaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}

