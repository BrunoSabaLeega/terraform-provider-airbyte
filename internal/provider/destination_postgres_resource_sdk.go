// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationPostgresResourceModel) ToCreateSDKType() *shared.DestinationPostgresCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationPostgresPostgres(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var sslMode *shared.DestinationPostgresSSLModes
	var destinationPostgresSSLModesDisable *shared.DestinationPostgresSSLModesDisable
	if r.Configuration.SslMode.DestinationPostgresSSLModesAllow != nil {
		mode := shared.DestinationPostgresSSLModesDisableMode(r.Configuration.SslMode.DestinationPostgresSSLModesAllow.Mode.ValueString())
		destinationPostgresSSLModesDisable = &shared.DestinationPostgresSSLModesDisable{
			Mode: mode,
		}
	}
	if destinationPostgresSSLModesDisable != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesDisable: destinationPostgresSSLModesDisable,
		}
	}
	var destinationPostgresSSLModesAllow *shared.DestinationPostgresSSLModesAllow
	if r.Configuration.SslMode.DestinationPostgresSSLModesDisable != nil {
		mode1 := shared.DestinationPostgresSSLModesAllowMode(r.Configuration.SslMode.DestinationPostgresSSLModesDisable.Mode.ValueString())
		destinationPostgresSSLModesAllow = &shared.DestinationPostgresSSLModesAllow{
			Mode: mode1,
		}
	}
	if destinationPostgresSSLModesAllow != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesAllow: destinationPostgresSSLModesAllow,
		}
	}
	var destinationPostgresSSLModesPrefer *shared.DestinationPostgresSSLModesPrefer
	if r.Configuration.SslMode.DestinationPostgresSSLModesPrefer != nil {
		mode2 := shared.DestinationPostgresSSLModesPreferMode(r.Configuration.SslMode.DestinationPostgresSSLModesPrefer.Mode.ValueString())
		destinationPostgresSSLModesPrefer = &shared.DestinationPostgresSSLModesPrefer{
			Mode: mode2,
		}
	}
	if destinationPostgresSSLModesPrefer != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesPrefer: destinationPostgresSSLModesPrefer,
		}
	}
	var destinationPostgresSSLModesRequire *shared.DestinationPostgresSSLModesRequire
	if r.Configuration.SslMode.DestinationPostgresSSLModesRequire != nil {
		mode3 := shared.DestinationPostgresSSLModesRequireMode(r.Configuration.SslMode.DestinationPostgresSSLModesRequire.Mode.ValueString())
		destinationPostgresSSLModesRequire = &shared.DestinationPostgresSSLModesRequire{
			Mode: mode3,
		}
	}
	if destinationPostgresSSLModesRequire != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesRequire: destinationPostgresSSLModesRequire,
		}
	}
	var destinationPostgresSSLModesVerifyCa *shared.DestinationPostgresSSLModesVerifyCa
	if r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa != nil {
		caCertificate := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.CaCertificate.ValueString()
		clientKeyPassword := new(string)
		if !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.IsNull() {
			*clientKeyPassword = r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword = nil
		}
		mode4 := shared.DestinationPostgresSSLModesVerifyCaMode(r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.Mode.ValueString())
		destinationPostgresSSLModesVerifyCa = &shared.DestinationPostgresSSLModesVerifyCa{
			CaCertificate:     caCertificate,
			ClientKeyPassword: clientKeyPassword,
			Mode:              mode4,
		}
	}
	if destinationPostgresSSLModesVerifyCa != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesVerifyCa: destinationPostgresSSLModesVerifyCa,
		}
	}
	var destinationPostgresSSLModesVerifyFull *shared.DestinationPostgresSSLModesVerifyFull
	if r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull != nil {
		caCertificate1 := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.CaCertificate.ValueString()
		clientCertificate := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientCertificate.ValueString()
		clientKey := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKey.ValueString()
		clientKeyPassword1 := new(string)
		if !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.IsNull() {
			*clientKeyPassword1 = r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword1 = nil
		}
		mode5 := shared.DestinationPostgresSSLModesVerifyFullMode(r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.Mode.ValueString())
		destinationPostgresSSLModesVerifyFull = &shared.DestinationPostgresSSLModesVerifyFull{
			CaCertificate:     caCertificate1,
			ClientCertificate: clientCertificate,
			ClientKey:         clientKey,
			ClientKeyPassword: clientKeyPassword1,
			Mode:              mode5,
		}
	}
	if destinationPostgresSSLModesVerifyFull != nil {
		sslMode = &shared.DestinationPostgresSSLModes{
			DestinationPostgresSSLModesVerifyFull: destinationPostgresSSLModesVerifyFull,
		}
	}
	var tunnelMethod *shared.DestinationPostgresSSHTunnelMethod
	var destinationPostgresSSHTunnelMethodNoTunnel *shared.DestinationPostgresSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationPostgresSSHTunnelMethodNoTunnel = &shared.DestinationPostgresSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationPostgresSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
			DestinationPostgresSSHTunnelMethodNoTunnel: destinationPostgresSSHTunnelMethodNoTunnel,
		}
	}
	var destinationPostgresSSHTunnelMethodSSHKeyAuthentication *shared.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationPostgresSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
			DestinationPostgresSSHTunnelMethodSSHKeyAuthentication: destinationPostgresSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationPostgresSSHTunnelMethodPasswordAuthentication *shared.DestinationPostgresSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationPostgresSSHTunnelMethodPasswordAuthentication = &shared.DestinationPostgresSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
			DestinationPostgresSSHTunnelMethodPasswordAuthentication: destinationPostgresSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationPostgres{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		Schema:          schema,
		SslMode:         sslMode,
		TunnelMethod:    tunnelMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPostgresCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPostgresResourceModel) ToUpdateSDKType() *shared.DestinationPostgresPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var sslMode *shared.DestinationPostgresUpdateSSLModes
	var destinationPostgresUpdateSSLModesDisable *shared.DestinationPostgresUpdateSSLModesDisable
	if r.Configuration.SslMode.DestinationPostgresSSLModesAllow != nil {
		mode := shared.DestinationPostgresUpdateSSLModesDisableMode(r.Configuration.SslMode.DestinationPostgresSSLModesAllow.Mode.ValueString())
		destinationPostgresUpdateSSLModesDisable = &shared.DestinationPostgresUpdateSSLModesDisable{
			Mode: mode,
		}
	}
	if destinationPostgresUpdateSSLModesDisable != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesDisable: destinationPostgresUpdateSSLModesDisable,
		}
	}
	var destinationPostgresUpdateSSLModesAllow *shared.DestinationPostgresUpdateSSLModesAllow
	if r.Configuration.SslMode.DestinationPostgresSSLModesDisable != nil {
		mode1 := shared.DestinationPostgresUpdateSSLModesAllowMode(r.Configuration.SslMode.DestinationPostgresSSLModesDisable.Mode.ValueString())
		destinationPostgresUpdateSSLModesAllow = &shared.DestinationPostgresUpdateSSLModesAllow{
			Mode: mode1,
		}
	}
	if destinationPostgresUpdateSSLModesAllow != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesAllow: destinationPostgresUpdateSSLModesAllow,
		}
	}
	var destinationPostgresUpdateSSLModesPrefer *shared.DestinationPostgresUpdateSSLModesPrefer
	if r.Configuration.SslMode.DestinationPostgresSSLModesPrefer != nil {
		mode2 := shared.DestinationPostgresUpdateSSLModesPreferMode(r.Configuration.SslMode.DestinationPostgresSSLModesPrefer.Mode.ValueString())
		destinationPostgresUpdateSSLModesPrefer = &shared.DestinationPostgresUpdateSSLModesPrefer{
			Mode: mode2,
		}
	}
	if destinationPostgresUpdateSSLModesPrefer != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesPrefer: destinationPostgresUpdateSSLModesPrefer,
		}
	}
	var destinationPostgresUpdateSSLModesRequire *shared.DestinationPostgresUpdateSSLModesRequire
	if r.Configuration.SslMode.DestinationPostgresSSLModesRequire != nil {
		mode3 := shared.DestinationPostgresUpdateSSLModesRequireMode(r.Configuration.SslMode.DestinationPostgresSSLModesRequire.Mode.ValueString())
		destinationPostgresUpdateSSLModesRequire = &shared.DestinationPostgresUpdateSSLModesRequire{
			Mode: mode3,
		}
	}
	if destinationPostgresUpdateSSLModesRequire != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesRequire: destinationPostgresUpdateSSLModesRequire,
		}
	}
	var destinationPostgresUpdateSSLModesVerifyCa *shared.DestinationPostgresUpdateSSLModesVerifyCa
	if r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa != nil {
		caCertificate := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.CaCertificate.ValueString()
		clientKeyPassword := new(string)
		if !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.IsNull() {
			*clientKeyPassword = r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword = nil
		}
		mode4 := shared.DestinationPostgresUpdateSSLModesVerifyCaMode(r.Configuration.SslMode.DestinationPostgresSSLModesVerifyCa.Mode.ValueString())
		destinationPostgresUpdateSSLModesVerifyCa = &shared.DestinationPostgresUpdateSSLModesVerifyCa{
			CaCertificate:     caCertificate,
			ClientKeyPassword: clientKeyPassword,
			Mode:              mode4,
		}
	}
	if destinationPostgresUpdateSSLModesVerifyCa != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesVerifyCa: destinationPostgresUpdateSSLModesVerifyCa,
		}
	}
	var destinationPostgresUpdateSSLModesVerifyFull *shared.DestinationPostgresUpdateSSLModesVerifyFull
	if r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull != nil {
		caCertificate1 := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.CaCertificate.ValueString()
		clientCertificate := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientCertificate.ValueString()
		clientKey := r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKey.ValueString()
		clientKeyPassword1 := new(string)
		if !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.IsNull() {
			*clientKeyPassword1 = r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.ClientKeyPassword.ValueString()
		} else {
			clientKeyPassword1 = nil
		}
		mode5 := shared.DestinationPostgresUpdateSSLModesVerifyFullMode(r.Configuration.SslMode.DestinationPostgresSSLModesVerifyFull.Mode.ValueString())
		destinationPostgresUpdateSSLModesVerifyFull = &shared.DestinationPostgresUpdateSSLModesVerifyFull{
			CaCertificate:     caCertificate1,
			ClientCertificate: clientCertificate,
			ClientKey:         clientKey,
			ClientKeyPassword: clientKeyPassword1,
			Mode:              mode5,
		}
	}
	if destinationPostgresUpdateSSLModesVerifyFull != nil {
		sslMode = &shared.DestinationPostgresUpdateSSLModes{
			DestinationPostgresUpdateSSLModesVerifyFull: destinationPostgresUpdateSSLModesVerifyFull,
		}
	}
	var tunnelMethod *shared.DestinationPostgresUpdateSSHTunnelMethod
	var destinationPostgresUpdateSSHTunnelMethodNoTunnel *shared.DestinationPostgresUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationPostgresUpdateSSHTunnelMethodNoTunnel = &shared.DestinationPostgresUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationPostgresUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
			DestinationPostgresUpdateSSHTunnelMethodNoTunnel: destinationPostgresUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
			DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication: destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
			DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication: destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationPostgresUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schema:        schema,
		SslMode:       sslMode,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPostgresPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPostgresResourceModel) ToDeleteSDKType() *shared.DestinationPostgresCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationPostgresResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
