// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DestinationMilvusDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationMilvusDataSource{}

func NewDestinationMilvusDataSource() datasource.DataSource {
	return &DestinationMilvusDataSource{}
}

// DestinationMilvusDataSource is the data source implementation.
type DestinationMilvusDataSource struct {
	client *sdk.SDK
}

// DestinationMilvusDataSourceModel describes the data model.
type DestinationMilvusDataSourceModel struct {
	Configuration DestinationMilvus `tfsdk:"configuration"`
	DestinationID types.String      `tfsdk:"destination_id"`
	Name          types.String      `tfsdk:"name"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationMilvusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_milvus"
}

// Schema defines the schema for the data source.
func (r *DestinationMilvusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationMilvus DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"milvus",
							),
						},
						Description: `must be one of ["milvus"]`,
					},
					"embedding": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_milvus_embedding_cohere": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"cohere_key": schema.StringAttribute{
										Computed: true,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"cohere",
											),
										},
										Description: `must be one of ["cohere"]`,
									},
								},
								Description: `Use the Cohere API to embed text.`,
							},
							"destination_milvus_embedding_fake": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"fake",
											),
										},
										Description: `must be one of ["fake"]`,
									},
								},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
							},
							"destination_milvus_embedding_from_field": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"dimensions": schema.Int64Attribute{
										Computed:    true,
										Description: `The number of dimensions the embedding model is generating`,
									},
									"field_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the field in the record that contains the embedding`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"from_field",
											),
										},
										Description: `must be one of ["from_field"]`,
									},
								},
								Description: `Use a field in the record as the embedding. This is useful if you already have an embedding for your data and want to store it in the vector store.`,
							},
							"destination_milvus_embedding_open_ai": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"openai",
											),
										},
										Description: `must be one of ["openai"]`,
									},
									"openai_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
							"destination_milvus_update_embedding_cohere": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"cohere_key": schema.StringAttribute{
										Computed: true,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"cohere",
											),
										},
										Description: `must be one of ["cohere"]`,
									},
								},
								Description: `Use the Cohere API to embed text.`,
							},
							"destination_milvus_update_embedding_fake": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"fake",
											),
										},
										Description: `must be one of ["fake"]`,
									},
								},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
							},
							"destination_milvus_update_embedding_from_field": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"dimensions": schema.Int64Attribute{
										Computed:    true,
										Description: `The number of dimensions the embedding model is generating`,
									},
									"field_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the field in the record that contains the embedding`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"from_field",
											),
										},
										Description: `must be one of ["from_field"]`,
									},
								},
								Description: `Use a field in the record as the embedding. This is useful if you already have an embedding for your data and want to store it in the vector store.`,
							},
							"destination_milvus_update_embedding_open_ai": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"openai",
											),
										},
										Description: `must be one of ["openai"]`,
									},
									"openai_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Embedding configuration`,
					},
					"indexing": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"destination_milvus_indexing_authentication_api_token": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"token",
													),
												},
												Description: `must be one of ["token"]`,
											},
											"token": schema.StringAttribute{
												Computed:    true,
												Description: `API Token for the Milvus instance`,
											},
										},
										Description: `Authenticate using an API token (suitable for Zilliz Cloud)`,
									},
									"destination_milvus_indexing_authentication_no_auth": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"no_auth",
													),
												},
												Description: `must be one of ["no_auth"]`,
											},
										},
										Description: `Do not authenticate (suitable for locally running test clusters, do not use for clusters with public IP addresses)`,
									},
									"destination_milvus_indexing_authentication_username_password": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"username_password",
													),
												},
												Description: `must be one of ["username_password"]`,
											},
											"password": schema.StringAttribute{
												Computed:    true,
												Description: `Password for the Milvus instance`,
											},
											"username": schema.StringAttribute{
												Computed:    true,
												Description: `Username for the Milvus instance`,
											},
										},
										Description: `Authenticate using username and password (suitable for self-managed Milvus clusters)`,
									},
									"destination_milvus_update_indexing_authentication_api_token": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"token",
													),
												},
												Description: `must be one of ["token"]`,
											},
											"token": schema.StringAttribute{
												Computed:    true,
												Description: `API Token for the Milvus instance`,
											},
										},
										Description: `Authenticate using an API token (suitable for Zilliz Cloud)`,
									},
									"destination_milvus_update_indexing_authentication_no_auth": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"no_auth",
													),
												},
												Description: `must be one of ["no_auth"]`,
											},
										},
										Description: `Do not authenticate (suitable for locally running test clusters, do not use for clusters with public IP addresses)`,
									},
									"destination_milvus_update_indexing_authentication_username_password": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"username_password",
													),
												},
												Description: `must be one of ["username_password"]`,
											},
											"password": schema.StringAttribute{
												Computed:    true,
												Description: `Password for the Milvus instance`,
											},
											"username": schema.StringAttribute{
												Computed:    true,
												Description: `Username for the Milvus instance`,
											},
										},
										Description: `Authenticate using username and password (suitable for self-managed Milvus clusters)`,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								Description: `Authentication method`,
							},
							"collection": schema.StringAttribute{
								Computed:    true,
								Description: `The collection to load data into`,
							},
							"db": schema.StringAttribute{
								Computed:    true,
								Description: `The database to connect to`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Description: `The public endpoint of the Milvus instance. `,
							},
							"text_field": schema.StringAttribute{
								Computed:    true,
								Description: `The field in the entity that contains the embedded text`,
							},
							"vector_field": schema.StringAttribute{
								Computed:    true,
								Description: `The field in the entity that contains the vector`,
							},
						},
						Description: `Indexing configuration`,
					},
					"processing": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"chunk_overlap": schema.Int64Attribute{
								Computed:    true,
								Description: `Size of overlap between chunks in tokens to store in vector store to better capture relevant context`,
							},
							"chunk_size": schema.Int64Attribute{
								Computed:    true,
								Description: `Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)`,
							},
							"metadata_fields": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.`,
							},
							"text_fields": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array.`,
							},
						},
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationMilvusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationMilvusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationMilvusDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationMilvusRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationMilvus(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
