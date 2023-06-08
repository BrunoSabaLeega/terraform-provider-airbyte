// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationS3GlueResource{}
var _ resource.ResourceWithImportState = &DestinationS3GlueResource{}

func NewDestinationS3GlueResource() resource.Resource {
	return &DestinationS3GlueResource{}
}

// DestinationS3GlueResource defines the resource implementation.
type DestinationS3GlueResource struct {
	client *sdk.SDK
}

// DestinationS3GlueResourceModel describes the resource data model.
type DestinationS3GlueResourceModel struct {
	Configuration   DestinationS3Glue `tfsdk:"configuration"`
	DestinationID   types.String      `tfsdk:"destination_id"`
	DestinationType types.String      `tfsdk:"destination_type"`
	Name            types.String      `tfsdk:"name"`
	WorkspaceID     types.String      `tfsdk:"workspace_id"`
}

func (r *DestinationS3GlueResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_s3_glue"
}

func (r *DestinationS3GlueResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationS3Glue Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Optional: true,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"s3-glue",
							),
						},
					},
					"file_name_pattern": schema.StringAttribute{
						Optional: true,
					},
					"format": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"destination_s3_glue_output_format_json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"destination_s3_glue_output_format_json_lines_newline_delimited_json_compression_no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
											"destination_s3_glue_output_format_json_lines_newline_delimited_json_compression_gzip": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"flattening": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
										Description: `Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.`,
									},
									"format_type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
									},
								},
								Description: `Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details`,
							},
							"destination_s3_glue_update_output_format_json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"destination_s3_glue_update_output_format_json_lines_newline_delimited_json_compression_no_compression": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
											"destination_s3_glue_update_output_format_json_lines_newline_delimited_json_compression_gzip": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Computed: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"flattening": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
										Description: `Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.`,
									},
									"format_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
									},
								},
								Description: `Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"glue_database": schema.StringAttribute{
						Required: true,
					},
					"glue_serialization_library": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"org.openx.data.jsonserde.JsonSerDe",
								"org.apache.hive.hcatalog.data.JsonSerDe",
							),
						},
						Description: `The library that your query engine will use for reading and writing data in your lake.`,
					},
					"s3_bucket_name": schema.StringAttribute{
						Required: true,
					},
					"s3_bucket_path": schema.StringAttribute{
						Required: true,
					},
					"s3_bucket_region": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"us-east-1",
								"us-east-2",
								"us-west-1",
								"us-west-2",
								"af-south-1",
								"ap-east-1",
								"ap-south-1",
								"ap-northeast-1",
								"ap-northeast-2",
								"ap-northeast-3",
								"ap-southeast-1",
								"ap-southeast-2",
								"ca-central-1",
								"cn-north-1",
								"cn-northwest-1",
								"eu-central-1",
								"eu-north-1",
								"eu-south-1",
								"eu-west-1",
								"eu-west-2",
								"eu-west-3",
								"sa-east-1",
								"me-south-1",
								"us-gov-east-1",
								"us-gov-west-1",
							),
						},
						Description: `The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.`,
					},
					"s3_endpoint": schema.StringAttribute{
						Optional: true,
					},
					"s3_path_format": schema.StringAttribute{
						Optional: true,
					},
					"secret_access_key": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DestinationS3GlueResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationS3GlueResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationS3GlueResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Destinations.CreateDestinationS3Glue(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3GlueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationS3GlueResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3GlueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationS3GlueResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationS3GluePutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationS3GlueRequest{
		DestinationS3GluePutRequest: destinationS3GluePutRequest,
		DestinationID:               destinationID,
	}
	res, err := r.client.Destinations.PutDestinationS3Glue(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationS3GlueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationS3GlueResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteDestinationS3GlueRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationS3Glue(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationS3GlueResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
