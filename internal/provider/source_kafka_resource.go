//source_kafka_resource

package provider

import (
	"airbyte/internal/sdk"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &SourceKafkaResource{}
var _ resource.ResourceWithImportState = &SourceKafkaResource{}

func NewSourceKafkaResource() resource.Resource {
	return &SourceKafkaResource{}
}

type SourceKafkaResource struct {
	client *sdk.SDK
}

func (r *SourceKafkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_kafka"
}