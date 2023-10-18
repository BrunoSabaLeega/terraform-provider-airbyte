//source_kafka_resource

package provider

import (
	"airbyte/internal/sdk"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &SourceKlarnaDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceKlarnaDataSource{}

func NewSourceKafkaDataSource() datasource.DataSource {
	return &SourceKafkaDataSource{}
}

// SourceKafkaDataSource is the data source implementation.
type SourceKafkaDataSource struct {
	client *sdk.SDK
}

func (r *SourceKafkaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_kafka"
}