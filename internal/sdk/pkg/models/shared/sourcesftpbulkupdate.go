// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// SourceSftpBulkUpdateFileType - The file type you want to sync. Currently only 'csv' and 'json' files are supported.
type SourceSftpBulkUpdateFileType string

const (
	SourceSftpBulkUpdateFileTypeCsv  SourceSftpBulkUpdateFileType = "csv"
	SourceSftpBulkUpdateFileTypeJSON SourceSftpBulkUpdateFileType = "json"
)

func (e SourceSftpBulkUpdateFileType) ToPointer() *SourceSftpBulkUpdateFileType {
	return &e
}

func (e *SourceSftpBulkUpdateFileType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		*e = SourceSftpBulkUpdateFileType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpBulkUpdateFileType: %v", v)
	}
}

type SourceSftpBulkUpdate struct {
	// Sync only the most recent file for the configured folder path and file pattern
	FileMostRecent *bool `json:"file_most_recent,omitempty"`
	// The regular expression to specify files for sync in a chosen Folder Path
	FilePattern *string `json:"file_pattern,omitempty"`
	// The file type you want to sync. Currently only 'csv' and 'json' files are supported.
	FileType *SourceSftpBulkUpdateFileType `json:"file_type,omitempty"`
	// The directory to search files for sync
	FolderPath string `json:"folder_path"`
	// The server host address
	Host string `json:"host"`
	// OS-level password for logging into the jump server host
	Password *string `json:"password,omitempty"`
	// The server port
	Port int64 `json:"port"`
	// The private key
	PrivateKey *string `json:"private_key,omitempty"`
	// The separator used in the CSV files. Define None if you want to use the Sniffer functionality
	Separator *string `json:"separator,omitempty"`
	// The date from which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// The name of the stream or table you want to create
	StreamName string `json:"stream_name"`
	// The server user
	Username string `json:"username"`
}