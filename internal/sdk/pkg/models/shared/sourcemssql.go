// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod string

const (
	SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethodStandard SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod = "STANDARD"
)

func (e SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod) ToPointer() *SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod {
	return &e
}

func (e *SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		*e = SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod: %v", v)
	}
}

// SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor - Incrementally detects new inserts and updates using the <a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor">cursor column</a> chosen when configuring a connection (e.g. created_at, updated_at).
type SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor struct {
	Method SourceMssqlUpdateMethodScanChangesWithUserDefinedCursorMethod `json:"method"`
}

// SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync - What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
type SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync string

const (
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSyncExistingAndNew SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync = "Existing and New"
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSyncNewChangesOnly SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync = "New Changes Only"
)

func (e SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync) ToPointer() *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync {
	return &e
}

func (e *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Existing and New":
		fallthrough
	case "New Changes Only":
		*e = SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync: %v", v)
	}
}

type SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod string

const (
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethodCdc SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod = "CDC"
)

func (e SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod) ToPointer() *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod {
	return &e
}

func (e *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CDC":
		*e = SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod: %v", v)
	}
}

// SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel - Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
type SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel string

const (
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevelSnapshot      SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel = "Snapshot"
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevelReadCommitted SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel = "Read Committed"
)

func (e SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel) ToPointer() *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel {
	return &e
}

func (e *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Snapshot":
		fallthrough
	case "Read Committed":
		*e = SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel: %v", v)
	}
}

// SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC - <i>Recommended</i> - Incrementally reads new inserts, updates, and deletes using the SQL Server's <a href="https://docs.airbyte.com/integrations/sources/mssql/#change-data-capture-cdc">change data capture feature</a>. This must be enabled on your database.
type SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC struct {
	// What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
	DataToSync *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCDataToSync `json:"data_to_sync,omitempty"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
	InitialWaitingSeconds *int64                                                            `json:"initial_waiting_seconds,omitempty"`
	Method                SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCMethod `json:"method"`
	// Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
	SnapshotIsolation *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDCInitialSnapshotIsolationLevel `json:"snapshot_isolation,omitempty"`
}

type SourceMssqlUpdateMethodType string

const (
	SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC SourceMssqlUpdateMethodType = "source-mssql_Update Method_Read Changes using Change Data Capture (CDC)"
	SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodScanChangesWithUserDefinedCursor     SourceMssqlUpdateMethodType = "source-mssql_Update Method_Scan Changes with User Defined Cursor"
)

type SourceMssqlUpdateMethod struct {
	SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC *SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC
	SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor     *SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor

	Type SourceMssqlUpdateMethodType
}

func CreateSourceMssqlUpdateMethodSourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC(sourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC) SourceMssqlUpdateMethod {
	typ := SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC

	return SourceMssqlUpdateMethod{
		SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC: &sourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC,
		Type: typ,
	}
}

func CreateSourceMssqlUpdateMethodSourceMssqlUpdateMethodScanChangesWithUserDefinedCursor(sourceMssqlUpdateMethodScanChangesWithUserDefinedCursor SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor) SourceMssqlUpdateMethod {
	typ := SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodScanChangesWithUserDefinedCursor

	return SourceMssqlUpdateMethod{
		SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor: &sourceMssqlUpdateMethodScanChangesWithUserDefinedCursor,
		Type: typ,
	}
}

func (u *SourceMssqlUpdateMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlUpdateMethodScanChangesWithUserDefinedCursor := new(SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateMethodScanChangesWithUserDefinedCursor); err == nil {
		u.SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor = sourceMssqlUpdateMethodScanChangesWithUserDefinedCursor
		u.Type = SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodScanChangesWithUserDefinedCursor
		return nil
	}

	sourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC := new(SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC); err == nil {
		u.SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC = sourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC
		u.Type = SourceMssqlUpdateMethodTypeSourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlUpdateMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor != nil {
		return json.Marshal(u.SourceMssqlUpdateMethodScanChangesWithUserDefinedCursor)
	}

	if u.SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC != nil {
		return json.Marshal(u.SourceMssqlUpdateMethodReadChangesUsingChangeDataCaptureCDC)
	}

	return nil, nil
}

type SourceMssqlMssql string

const (
	SourceMssqlMssqlMssql SourceMssqlMssql = "mssql"
)

func (e SourceMssqlMssql) ToPointer() *SourceMssqlMssql {
	return &e
}

func (e *SourceMssqlMssql) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mssql":
		*e = SourceMssqlMssql(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlMssql: %v", v)
	}
}

type SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod string

const (
	SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEncryptedVerifyCertificate SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod = "encrypted_verify_certificate"
)

func (e SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod) ToPointer() *SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod {
	return &e
}

func (e *SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod: %v", v)
	}
}

// SourceMssqlSSLMethodEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type SourceMssqlSSLMethodEncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                                 `json:"hostNameInCertificate,omitempty"`
	SslMethod             SourceMssqlSSLMethodEncryptedVerifyCertificateSSLMethod `json:"ssl_method"`
}

type SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod string

const (
	SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEncryptedTrustServerCertificate SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod = "encrypted_trust_server_certificate"
)

func (e SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod) ToPointer() *SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod {
	return &e
}

func (e *SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_trust_server_certificate":
		*e = SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod: %v", v)
	}
}

// SourceMssqlSSLMethodEncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type SourceMssqlSSLMethodEncryptedTrustServerCertificate struct {
	SslMethod SourceMssqlSSLMethodEncryptedTrustServerCertificateSSLMethod `json:"ssl_method"`
}

type SourceMssqlSSLMethodType string

const (
	SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate SourceMssqlSSLMethodType = "source-mssql_SSL Method_Encrypted (trust server certificate)"
	SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate      SourceMssqlSSLMethodType = "source-mssql_SSL Method_Encrypted (verify certificate)"
)

type SourceMssqlSSLMethod struct {
	SourceMssqlSSLMethodEncryptedTrustServerCertificate *SourceMssqlSSLMethodEncryptedTrustServerCertificate
	SourceMssqlSSLMethodEncryptedVerifyCertificate      *SourceMssqlSSLMethodEncryptedVerifyCertificate

	Type SourceMssqlSSLMethodType
}

func CreateSourceMssqlSSLMethodSourceMssqlSSLMethodEncryptedTrustServerCertificate(sourceMssqlSSLMethodEncryptedTrustServerCertificate SourceMssqlSSLMethodEncryptedTrustServerCertificate) SourceMssqlSSLMethod {
	typ := SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate

	return SourceMssqlSSLMethod{
		SourceMssqlSSLMethodEncryptedTrustServerCertificate: &sourceMssqlSSLMethodEncryptedTrustServerCertificate,
		Type: typ,
	}
}

func CreateSourceMssqlSSLMethodSourceMssqlSSLMethodEncryptedVerifyCertificate(sourceMssqlSSLMethodEncryptedVerifyCertificate SourceMssqlSSLMethodEncryptedVerifyCertificate) SourceMssqlSSLMethod {
	typ := SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate

	return SourceMssqlSSLMethod{
		SourceMssqlSSLMethodEncryptedVerifyCertificate: &sourceMssqlSSLMethodEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *SourceMssqlSSLMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlSSLMethodEncryptedTrustServerCertificate := new(SourceMssqlSSLMethodEncryptedTrustServerCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSLMethodEncryptedTrustServerCertificate); err == nil {
		u.SourceMssqlSSLMethodEncryptedTrustServerCertificate = sourceMssqlSSLMethodEncryptedTrustServerCertificate
		u.Type = SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedTrustServerCertificate
		return nil
	}

	sourceMssqlSSLMethodEncryptedVerifyCertificate := new(SourceMssqlSSLMethodEncryptedVerifyCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSLMethodEncryptedVerifyCertificate); err == nil {
		u.SourceMssqlSSLMethodEncryptedVerifyCertificate = sourceMssqlSSLMethodEncryptedVerifyCertificate
		u.Type = SourceMssqlSSLMethodTypeSourceMssqlSSLMethodEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlSSLMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlSSLMethodEncryptedTrustServerCertificate != nil {
		return json.Marshal(u.SourceMssqlSSLMethodEncryptedTrustServerCertificate)
	}

	if u.SourceMssqlSSLMethodEncryptedVerifyCertificate != nil {
		return json.Marshal(u.SourceMssqlSSLMethodEncryptedVerifyCertificate)
	}

	return nil, nil
}

// SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMssqlSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod SourceMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMssqlSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod SourceMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod string

const (
	SourceMssqlSSHTunnelMethodNoTunnelTunnelMethodNoTunnel SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// SourceMssqlSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod SourceMssqlSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type SourceMssqlSSHTunnelMethodType string

const (
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel               SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_No Tunnel"
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication   SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_SSH Key Authentication"
	SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication SourceMssqlSSHTunnelMethodType = "source-mssql_SSH Tunnel Method_Password Authentication"
)

type SourceMssqlSSHTunnelMethod struct {
	SourceMssqlSSHTunnelMethodNoTunnel               *SourceMssqlSSHTunnelMethodNoTunnel
	SourceMssqlSSHTunnelMethodSSHKeyAuthentication   *SourceMssqlSSHTunnelMethodSSHKeyAuthentication
	SourceMssqlSSHTunnelMethodPasswordAuthentication *SourceMssqlSSHTunnelMethodPasswordAuthentication

	Type SourceMssqlSSHTunnelMethodType
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodNoTunnel(sourceMssqlSSHTunnelMethodNoTunnel SourceMssqlSSHTunnelMethodNoTunnel) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodNoTunnel: &sourceMssqlSSHTunnelMethodNoTunnel,
		Type:                               typ,
	}
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodSSHKeyAuthentication(sourceMssqlSSHTunnelMethodSSHKeyAuthentication SourceMssqlSSHTunnelMethodSSHKeyAuthentication) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodSSHKeyAuthentication: &sourceMssqlSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateSourceMssqlSSHTunnelMethodSourceMssqlSSHTunnelMethodPasswordAuthentication(sourceMssqlSSHTunnelMethodPasswordAuthentication SourceMssqlSSHTunnelMethodPasswordAuthentication) SourceMssqlSSHTunnelMethod {
	typ := SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication

	return SourceMssqlSSHTunnelMethod{
		SourceMssqlSSHTunnelMethodPasswordAuthentication: &sourceMssqlSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *SourceMssqlSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlSSHTunnelMethodNoTunnel := new(SourceMssqlSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodNoTunnel); err == nil {
		u.SourceMssqlSSHTunnelMethodNoTunnel = sourceMssqlSSHTunnelMethodNoTunnel
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodNoTunnel
		return nil
	}

	sourceMssqlSSHTunnelMethodSSHKeyAuthentication := new(SourceMssqlSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication = sourceMssqlSSHTunnelMethodSSHKeyAuthentication
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	sourceMssqlSSHTunnelMethodPasswordAuthentication := new(SourceMssqlSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlSSHTunnelMethodPasswordAuthentication); err == nil {
		u.SourceMssqlSSHTunnelMethodPasswordAuthentication = sourceMssqlSSHTunnelMethodPasswordAuthentication
		u.Type = SourceMssqlSSHTunnelMethodTypeSourceMssqlSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodNoTunnel)
	}

	if u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.SourceMssqlSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.SourceMssqlSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type SourceMssql struct {
	// The name of the database.
	Database string `json:"database"`
	// The hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port of the database.
	Port int64 `json:"port"`
	// Configures how data is extracted from the database.
	ReplicationMethod *SourceMssqlUpdateMethod `json:"replication_method,omitempty"`
	// The list of schemas to sync from. Defaults to user. Case sensitive.
	Schemas    []string         `json:"schemas,omitempty"`
	SourceType SourceMssqlMssql `json:"sourceType"`
	// The encryption method which is used when communicating with the database.
	SslMethod *SourceMssqlSSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceMssqlSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}
