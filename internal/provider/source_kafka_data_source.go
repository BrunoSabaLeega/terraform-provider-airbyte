//source_kafka_data_source

package provider

import (
	"airbyte/internal/sdk"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceKafkaDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceKafkaDataSource{}

func NewSourceKafkaDataSource() datasource.DataSource {
	return &SourceKafkaDataSource{}
}

// SourceKafkaDataSource is the data source implementation.
type SourceKafkaDataSource struct {
	client *sdk.SDK
}
// SourceKafkaDataSourceModel describes the data model.
type SourceKafkaDataSourceModel struct {
	Configuration SourceKafka `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceKafkaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_kafka"
}