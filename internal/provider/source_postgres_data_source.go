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
var _ datasource.DataSource = &SourcePostgresDataSource{}
var _ datasource.DataSourceWithConfigure = &SourcePostgresDataSource{}

func NewSourcePostgresDataSource() datasource.DataSource {
	return &SourcePostgresDataSource{}
}

// SourcePostgresDataSource is the data source implementation.
type SourcePostgresDataSource struct {
	client *sdk.SDK
}

// SourcePostgresDataSourceModel describes the data model.
type SourcePostgresDataSourceModel struct {
	Configuration SourcePostgres1 `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourcePostgresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_postgres"
}

// Schema defines the schema for the data source.
func (r *SourcePostgresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourcePostgres DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the database.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `Hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (Eg. key1=value1&key2=value2&key3=value3). For more information read about <a href="https://jdbc.postgresql.org/documentation/head/connect.html">JDBC URL parameters</a>.`,
					},
					"password": schema.StringAttribute{
						Computed:    true,
						Description: `Password associated with the username.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `Port of the database.`,
					},
					"replication_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed:    true,
										Description: `The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-5-optional-set-up-initial-waiting-time">initial waiting time</a>.`,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										MarkdownDescription: `must be one of ["While reading Data", "After loading Data in the destination"]` + "\n" +
											`Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
										Description: `must be one of ["CDC"]`,
									},
									"plugin": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										MarkdownDescription: `must be one of ["pgoutput"]` + "\n" +
											`A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Computed:    true,
										Description: `A Postgres publication used for consuming changes. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-4-create-publications-and-replication-identities-for-tables">publications and replication identities</a>.`,
									},
									"queue_size": schema.Int64Attribute{
										Computed:    true,
										Description: `The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.`,
									},
									"replication_slot": schema.StringAttribute{
										Computed:    true,
										Description: `A plugin logical replication slot. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-3-create-replication-slot">replication slots</a>.`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_postgres_replication_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
										Description: `must be one of ["Standard"]`,
									},
								},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_postgres_update_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed:    true,
										Description: `The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-5-optional-set-up-initial-waiting-time">initial waiting time</a>.`,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										MarkdownDescription: `must be one of ["While reading Data", "After loading Data in the destination"]` + "\n" +
											`Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
										Description: `must be one of ["CDC"]`,
									},
									"plugin": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										MarkdownDescription: `must be one of ["pgoutput"]` + "\n" +
											`A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Computed:    true,
										Description: `A Postgres publication used for consuming changes. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-4-create-publications-and-replication-identities-for-tables">publications and replication identities</a>.`,
									},
									"queue_size": schema.Int64Attribute{
										Computed:    true,
										Description: `The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.`,
									},
									"replication_slot": schema.StringAttribute{
										Computed:    true,
										Description: `A plugin logical replication slot. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-3-create-replication-slot">replication slots</a>.`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_postgres_update_replication_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
										Description: `must be one of ["Standard"]`,
									},
								},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Replication method for extracting data from the database.`,
					},
					"schemas": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The list of schemas (case sensitive) to sync from. Defaults to public.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"postgres",
							),
						},
						Description: `must be one of ["postgres"]`,
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_ssl_modes_allow": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"allow",
											),
										},
										Description: `must be one of ["allow"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_postgres_ssl_modes_disable": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
										Description: `must be one of ["disable"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_postgres_ssl_modes_prefer": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"prefer",
											),
										},
										Description: `must be one of ["prefer"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_postgres_ssl_modes_require": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"require",
											),
										},
										Description: `must be one of ["require"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_postgres_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Computed:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-ca",
											),
										},
										Description: `must be one of ["verify-ca"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_postgres_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Computed:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
										Description: `must be one of ["verify-full"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
							"source_postgres_update_ssl_modes_allow": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"allow",
											),
										},
										Description: `must be one of ["allow"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_postgres_update_ssl_modes_disable": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
										Description: `must be one of ["disable"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_postgres_update_ssl_modes_prefer": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"prefer",
											),
										},
										Description: `must be one of ["prefer"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_postgres_update_ssl_modes_require": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"require",
											),
										},
										Description: `must be one of ["require"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_postgres_update_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Computed:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-ca",
											),
										},
										Description: `must be one of ["verify-ca"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_postgres_update_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Computed:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Computed:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Computed:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
										Description: `must be one of ["verify-full"]`,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						MarkdownDescription: `SSL connection modes. ` + "\n" +
							`  Read more <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.`,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_postgres_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_postgres_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_postgres_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_postgres_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_postgres_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Computed:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Computed:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
					},
					"username": schema.StringAttribute{
						Computed:    true,
						Description: `Username to access the database.`,
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

func (r *SourcePostgresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourcePostgresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourcePostgresDataSourceModel
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
	request := operations.GetSourcePostgresRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourcePostgres(ctx, request)
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
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}