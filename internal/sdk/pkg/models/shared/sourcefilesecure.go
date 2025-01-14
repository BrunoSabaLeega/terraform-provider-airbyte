// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceFileSecureFileFormat - The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
type SourceFileSecureFileFormat string

const (
	SourceFileSecureFileFormatCsv         SourceFileSecureFileFormat = "csv"
	SourceFileSecureFileFormatJSON        SourceFileSecureFileFormat = "json"
	SourceFileSecureFileFormatJsonl       SourceFileSecureFileFormat = "jsonl"
	SourceFileSecureFileFormatExcel       SourceFileSecureFileFormat = "excel"
	SourceFileSecureFileFormatExcelBinary SourceFileSecureFileFormat = "excel_binary"
	SourceFileSecureFileFormatFeather     SourceFileSecureFileFormat = "feather"
	SourceFileSecureFileFormatParquet     SourceFileSecureFileFormat = "parquet"
	SourceFileSecureFileFormatYaml        SourceFileSecureFileFormat = "yaml"
)

func (e SourceFileSecureFileFormat) ToPointer() *SourceFileSecureFileFormat {
	return &e
}

func (e *SourceFileSecureFileFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		fallthrough
	case "jsonl":
		fallthrough
	case "excel":
		fallthrough
	case "excel_binary":
		fallthrough
	case "feather":
		fallthrough
	case "parquet":
		fallthrough
	case "yaml":
		*e = SourceFileSecureFileFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureFileFormat: %v", v)
	}
}

type SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage string

const (
	SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorageSftp SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage = "SFTP"
)

func (e SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage) ToPointer() *SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SFTP":
		*e = SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol struct {
	Host     string                                                               `json:"host"`
	Password *string                                                              `json:"password,omitempty"`
	Port     *string                                                              `json:"port,omitempty"`
	Storage  SourceFileSecureStorageProviderSFTPSecureFileTransferProtocolStorage `json:"storage"`
	User     string                                                               `json:"user"`
}

type SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage string

const (
	SourceFileSecureStorageProviderSCPSecureCopyProtocolStorageScp SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage = "SCP"
)

func (e SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage) ToPointer() *SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SCP":
		*e = SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderSCPSecureCopyProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderSCPSecureCopyProtocol struct {
	Host     string                                                      `json:"host"`
	Password *string                                                     `json:"password,omitempty"`
	Port     *string                                                     `json:"port,omitempty"`
	Storage  SourceFileSecureStorageProviderSCPSecureCopyProtocolStorage `json:"storage"`
	User     string                                                      `json:"user"`
}

type SourceFileSecureStorageProviderSSHSecureShellStorage string

const (
	SourceFileSecureStorageProviderSSHSecureShellStorageSSH SourceFileSecureStorageProviderSSHSecureShellStorage = "SSH"
)

func (e SourceFileSecureStorageProviderSSHSecureShellStorage) ToPointer() *SourceFileSecureStorageProviderSSHSecureShellStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderSSHSecureShellStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH":
		*e = SourceFileSecureStorageProviderSSHSecureShellStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderSSHSecureShellStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderSSHSecureShell - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderSSHSecureShell struct {
	Host     string                                               `json:"host"`
	Password *string                                              `json:"password,omitempty"`
	Port     *string                                              `json:"port,omitempty"`
	Storage  SourceFileSecureStorageProviderSSHSecureShellStorage `json:"storage"`
	User     string                                               `json:"user"`
}

type SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage string

const (
	SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorageAzBlob SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage = "AzBlob"
)

func (e SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage) ToPointer() *SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AzBlob":
		*e = SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderAzBlobAzureBlobStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderAzBlobAzureBlobStorage struct {
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
	SasToken *string `json:"sas_token,omitempty"`
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
	SharedKey *string                                                      `json:"shared_key,omitempty"`
	Storage   SourceFileSecureStorageProviderAzBlobAzureBlobStorageStorage `json:"storage"`
	// The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.
	StorageAccount string `json:"storage_account"`
}

type SourceFileSecureStorageProviderS3AmazonWebServicesStorage string

const (
	SourceFileSecureStorageProviderS3AmazonWebServicesStorageS3 SourceFileSecureStorageProviderS3AmazonWebServicesStorage = "S3"
)

func (e SourceFileSecureStorageProviderS3AmazonWebServicesStorage) ToPointer() *SourceFileSecureStorageProviderS3AmazonWebServicesStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderS3AmazonWebServicesStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = SourceFileSecureStorageProviderS3AmazonWebServicesStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderS3AmazonWebServicesStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderS3AmazonWebServices - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderS3AmazonWebServices struct {
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsSecretAccessKey *string                                                   `json:"aws_secret_access_key,omitempty"`
	Storage            SourceFileSecureStorageProviderS3AmazonWebServicesStorage `json:"storage"`
}

type SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage string

const (
	SourceFileSecureStorageProviderGCSGoogleCloudStorageStorageGcs SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage = "GCS"
)

func (e SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage) ToPointer() *SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS":
		*e = SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderGCSGoogleCloudStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderGCSGoogleCloudStorage struct {
	// In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
	ServiceAccountJSON *string                                                     `json:"service_account_json,omitempty"`
	Storage            SourceFileSecureStorageProviderGCSGoogleCloudStorageStorage `json:"storage"`
}

type SourceFileSecureStorageProviderHTTPSPublicWebStorage string

const (
	SourceFileSecureStorageProviderHTTPSPublicWebStorageHTTPS SourceFileSecureStorageProviderHTTPSPublicWebStorage = "HTTPS"
)

func (e SourceFileSecureStorageProviderHTTPSPublicWebStorage) ToPointer() *SourceFileSecureStorageProviderHTTPSPublicWebStorage {
	return &e
}

func (e *SourceFileSecureStorageProviderHTTPSPublicWebStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTPS":
		*e = SourceFileSecureStorageProviderHTTPSPublicWebStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureStorageProviderHTTPSPublicWebStorage: %v", v)
	}
}

// SourceFileSecureStorageProviderHTTPSPublicWeb - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSecureStorageProviderHTTPSPublicWeb struct {
	Storage SourceFileSecureStorageProviderHTTPSPublicWebStorage `json:"storage"`
	// Add User-Agent to request
	UserAgent *bool `json:"user_agent,omitempty"`
}

type SourceFileSecureStorageProviderType string

const (
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderHTTPSPublicWeb                 SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_HTTPS: Public Web"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderGCSGoogleCloudStorage          SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_GCS: Google Cloud Storage"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderS3AmazonWebServices            SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_S3: Amazon Web Services"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderAzBlobAzureBlobStorage         SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_AzBlob: Azure Blob Storage"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSSHSecureShell                 SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_SSH: Secure Shell"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSCPSecureCopyProtocol          SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_SCP: Secure copy protocol"
	SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSFTPSecureFileTransferProtocol SourceFileSecureStorageProviderType = "source-file-secure_Storage Provider_SFTP: Secure File Transfer Protocol"
)

type SourceFileSecureStorageProvider struct {
	SourceFileSecureStorageProviderHTTPSPublicWeb                 *SourceFileSecureStorageProviderHTTPSPublicWeb
	SourceFileSecureStorageProviderGCSGoogleCloudStorage          *SourceFileSecureStorageProviderGCSGoogleCloudStorage
	SourceFileSecureStorageProviderS3AmazonWebServices            *SourceFileSecureStorageProviderS3AmazonWebServices
	SourceFileSecureStorageProviderAzBlobAzureBlobStorage         *SourceFileSecureStorageProviderAzBlobAzureBlobStorage
	SourceFileSecureStorageProviderSSHSecureShell                 *SourceFileSecureStorageProviderSSHSecureShell
	SourceFileSecureStorageProviderSCPSecureCopyProtocol          *SourceFileSecureStorageProviderSCPSecureCopyProtocol
	SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol *SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol

	Type SourceFileSecureStorageProviderType
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderHTTPSPublicWeb(sourceFileSecureStorageProviderHTTPSPublicWeb SourceFileSecureStorageProviderHTTPSPublicWeb) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderHTTPSPublicWeb

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderHTTPSPublicWeb: &sourceFileSecureStorageProviderHTTPSPublicWeb,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderGCSGoogleCloudStorage(sourceFileSecureStorageProviderGCSGoogleCloudStorage SourceFileSecureStorageProviderGCSGoogleCloudStorage) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderGCSGoogleCloudStorage

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderGCSGoogleCloudStorage: &sourceFileSecureStorageProviderGCSGoogleCloudStorage,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderS3AmazonWebServices(sourceFileSecureStorageProviderS3AmazonWebServices SourceFileSecureStorageProviderS3AmazonWebServices) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderS3AmazonWebServices

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderS3AmazonWebServices: &sourceFileSecureStorageProviderS3AmazonWebServices,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderAzBlobAzureBlobStorage(sourceFileSecureStorageProviderAzBlobAzureBlobStorage SourceFileSecureStorageProviderAzBlobAzureBlobStorage) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderAzBlobAzureBlobStorage

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderAzBlobAzureBlobStorage: &sourceFileSecureStorageProviderAzBlobAzureBlobStorage,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderSSHSecureShell(sourceFileSecureStorageProviderSSHSecureShell SourceFileSecureStorageProviderSSHSecureShell) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSSHSecureShell

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderSSHSecureShell: &sourceFileSecureStorageProviderSSHSecureShell,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderSCPSecureCopyProtocol(sourceFileSecureStorageProviderSCPSecureCopyProtocol SourceFileSecureStorageProviderSCPSecureCopyProtocol) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSCPSecureCopyProtocol

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderSCPSecureCopyProtocol: &sourceFileSecureStorageProviderSCPSecureCopyProtocol,
		Type: typ,
	}
}

func CreateSourceFileSecureStorageProviderSourceFileSecureStorageProviderSFTPSecureFileTransferProtocol(sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol) SourceFileSecureStorageProvider {
	typ := SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSFTPSecureFileTransferProtocol

	return SourceFileSecureStorageProvider{
		SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol: &sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol,
		Type: typ,
	}
}

func (u *SourceFileSecureStorageProvider) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceFileSecureStorageProviderHTTPSPublicWeb := new(SourceFileSecureStorageProviderHTTPSPublicWeb)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderHTTPSPublicWeb); err == nil {
		u.SourceFileSecureStorageProviderHTTPSPublicWeb = sourceFileSecureStorageProviderHTTPSPublicWeb
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderHTTPSPublicWeb
		return nil
	}

	sourceFileSecureStorageProviderGCSGoogleCloudStorage := new(SourceFileSecureStorageProviderGCSGoogleCloudStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderGCSGoogleCloudStorage); err == nil {
		u.SourceFileSecureStorageProviderGCSGoogleCloudStorage = sourceFileSecureStorageProviderGCSGoogleCloudStorage
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderGCSGoogleCloudStorage
		return nil
	}

	sourceFileSecureStorageProviderS3AmazonWebServices := new(SourceFileSecureStorageProviderS3AmazonWebServices)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderS3AmazonWebServices); err == nil {
		u.SourceFileSecureStorageProviderS3AmazonWebServices = sourceFileSecureStorageProviderS3AmazonWebServices
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderS3AmazonWebServices
		return nil
	}

	sourceFileSecureStorageProviderAzBlobAzureBlobStorage := new(SourceFileSecureStorageProviderAzBlobAzureBlobStorage)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderAzBlobAzureBlobStorage); err == nil {
		u.SourceFileSecureStorageProviderAzBlobAzureBlobStorage = sourceFileSecureStorageProviderAzBlobAzureBlobStorage
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderAzBlobAzureBlobStorage
		return nil
	}

	sourceFileSecureStorageProviderSSHSecureShell := new(SourceFileSecureStorageProviderSSHSecureShell)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderSSHSecureShell); err == nil {
		u.SourceFileSecureStorageProviderSSHSecureShell = sourceFileSecureStorageProviderSSHSecureShell
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSSHSecureShell
		return nil
	}

	sourceFileSecureStorageProviderSCPSecureCopyProtocol := new(SourceFileSecureStorageProviderSCPSecureCopyProtocol)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderSCPSecureCopyProtocol); err == nil {
		u.SourceFileSecureStorageProviderSCPSecureCopyProtocol = sourceFileSecureStorageProviderSCPSecureCopyProtocol
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSCPSecureCopyProtocol
		return nil
	}

	sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol := new(SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol); err == nil {
		u.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol = sourceFileSecureStorageProviderSFTPSecureFileTransferProtocol
		u.Type = SourceFileSecureStorageProviderTypeSourceFileSecureStorageProviderSFTPSecureFileTransferProtocol
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFileSecureStorageProvider) MarshalJSON() ([]byte, error) {
	if u.SourceFileSecureStorageProviderHTTPSPublicWeb != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderHTTPSPublicWeb)
	}

	if u.SourceFileSecureStorageProviderGCSGoogleCloudStorage != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderGCSGoogleCloudStorage)
	}

	if u.SourceFileSecureStorageProviderS3AmazonWebServices != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderS3AmazonWebServices)
	}

	if u.SourceFileSecureStorageProviderAzBlobAzureBlobStorage != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderAzBlobAzureBlobStorage)
	}

	if u.SourceFileSecureStorageProviderSSHSecureShell != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderSSHSecureShell)
	}

	if u.SourceFileSecureStorageProviderSCPSecureCopyProtocol != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderSCPSecureCopyProtocol)
	}

	if u.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol != nil {
		return json.Marshal(u.SourceFileSecureStorageProviderSFTPSecureFileTransferProtocol)
	}

	return nil, nil
}

type SourceFileSecureFileSecure string

const (
	SourceFileSecureFileSecureFileSecure SourceFileSecureFileSecure = "file-secure"
)

func (e SourceFileSecureFileSecure) ToPointer() *SourceFileSecureFileSecure {
	return &e
}

func (e *SourceFileSecureFileSecure) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "file-secure":
		*e = SourceFileSecureFileSecure(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSecureFileSecure: %v", v)
	}
}

type SourceFileSecure struct {
	// The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
	DatasetName string `json:"dataset_name"`
	// The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
	Format SourceFileSecureFileFormat `json:"format"`
	// The storage Provider or Location of the file(s) which should be replicated.
	Provider SourceFileSecureStorageProvider `json:"provider"`
	// This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.
	ReaderOptions *string                    `json:"reader_options,omitempty"`
	SourceType    SourceFileSecureFileSecure `json:"sourceType"`
	// The URL path to access the file which should be replicated.
	URL string `json:"url"`
}
