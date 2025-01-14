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
var _ datasource.DataSource = &SourceLinkedinAdsDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceLinkedinAdsDataSource{}

func NewSourceLinkedinAdsDataSource() datasource.DataSource {
	return &SourceLinkedinAdsDataSource{}
}

// SourceLinkedinAdsDataSource is the data source implementation.
type SourceLinkedinAdsDataSource struct {
	client *sdk.SDK
}

// SourceLinkedinAdsDataSourceModel describes the data model.
type SourceLinkedinAdsDataSourceModel struct {
	Configuration SourceLinkedinAds `tfsdk:"configuration"`
	Name          types.String      `tfsdk:"name"`
	SecretID      types.String      `tfsdk:"secret_id"`
	SourceID      types.String      `tfsdk:"source_id"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceLinkedinAdsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_linkedin_ads"
}

// Schema defines the schema for the data source.
func (r *SourceLinkedinAdsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceLinkedinAds DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"account_ids": schema.ListAttribute{
						Computed:    true,
						ElementType: types.Int64Type,
						Description: `Specify the account IDs to pull data from, separated by a space. Leave this field empty if you want to pull the data from all accounts accessible by the authenticated user. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn docs</a> to locate these IDs.`,
					},
					"ad_analytics_reports": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `The name for the custom report.`,
								},
								"pivot_by": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"COMPANY",
											"ACCOUNT",
											"SHARE",
											"CAMPAIGN",
											"CREATIVE",
											"CAMPAIGN_GROUP",
											"CONVERSION",
											"CONVERSATION_NODE",
											"CONVERSATION_NODE_OPTION_INDEX",
											"SERVING_LOCATION",
											"CARD_INDEX",
											"MEMBER_COMPANY_SIZE",
											"MEMBER_INDUSTRY",
											"MEMBER_SENIORITY",
											"MEMBER_JOB_TITLE ",
											"MEMBER_JOB_FUNCTION ",
											"MEMBER_COUNTRY_V2 ",
											"MEMBER_REGION_V2",
											"MEMBER_COMPANY",
											"PLACEMENT_NAME",
											"IMPRESSION_DEVICE_TYPE",
										),
									},
									MarkdownDescription: `must be one of ["COMPANY", "ACCOUNT", "SHARE", "CAMPAIGN", "CREATIVE", "CAMPAIGN_GROUP", "CONVERSION", "CONVERSATION_NODE", "CONVERSATION_NODE_OPTION_INDEX", "SERVING_LOCATION", "CARD_INDEX", "MEMBER_COMPANY_SIZE", "MEMBER_INDUSTRY", "MEMBER_SENIORITY", "MEMBER_JOB_TITLE ", "MEMBER_JOB_FUNCTION ", "MEMBER_COUNTRY_V2 ", "MEMBER_REGION_V2", "MEMBER_COMPANY", "PLACEMENT_NAME", "IMPRESSION_DEVICE_TYPE"]` + "\n" +
										`Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives.`,
								},
								"time_granularity": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"ALL",
											"DAILY",
											"MONTHLY",
											"YEARLY",
										),
									},
									MarkdownDescription: `must be one of ["ALL", "DAILY", "MONTHLY", "YEARLY"]` + "\n" +
										`Choose how to group the data in your report by time. The options are:<br>- 'ALL': A single result summarizing the entire time range.<br>- 'DAILY': Group results by each day.<br>- 'MONTHLY': Group results by each month.<br>- 'YEARLY': Group results by each year.<br>Selecting a time grouping helps you analyze trends and patterns over different time periods.`,
								},
							},
						},
					},
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_linkedin_ads_authentication_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The access token generated for your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_ads_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The client ID of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The key to refresh the expired access token. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
								},
							},
							"source_linkedin_ads_update_authentication_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `The access token generated for your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"access_token",
											),
										},
										Description: `must be one of ["access_token"]`,
									},
								},
							},
							"source_linkedin_ads_update_authentication_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"oAuth2.0",
											),
										},
										Description: `must be one of ["oAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The client ID of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `The key to refresh the expired access token. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.`,
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"linkedin-ads",
							),
						},
						Description: `must be one of ["linkedin-ads"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceLinkedinAdsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceLinkedinAdsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceLinkedinAdsDataSourceModel
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

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceLinkedinAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceLinkedinAds(ctx, request)
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
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
