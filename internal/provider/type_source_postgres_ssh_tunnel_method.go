// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourcePostgresSSHTunnelMethod struct {
	SourcePostgresSSHTunnelMethodNoTunnel               *DestinationClickhouseSSHTunnelMethodNoTunnel               `tfsdk:"source_postgres_ssh_tunnel_method_no_tunnel"`
	SourcePostgresSSHTunnelMethodSSHKeyAuthentication   *DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication   `tfsdk:"source_postgres_ssh_tunnel_method_ssh_key_authentication"`
	SourcePostgresSSHTunnelMethodPasswordAuthentication *DestinationClickhouseSSHTunnelMethodPasswordAuthentication `tfsdk:"source_postgres_ssh_tunnel_method_password_authentication"`
}