// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationSftpJSONSftpJSONEnum string

const (
	DestinationSftpJSONSftpJSONEnumSftpJSON DestinationSftpJSONSftpJSONEnum = "sftp-json"
)

func (e *DestinationSftpJSONSftpJSONEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "sftp-json":
		*e = DestinationSftpJSONSftpJSONEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSftpJSONSftpJSONEnum: %s", s)
	}
}

// DestinationSftpJSON - The values required to configure the destination.
type DestinationSftpJSON struct {
	DestinationType DestinationSftpJSONSftpJSONEnum `json:"destinationType"`
	// Path to the directory where json files will be written.
	DestinationPath string `json:"destination_path"`
	// Hostname of the SFTP server.
	Host string `json:"host"`
	// Password associated with the username.
	Password string `json:"password"`
	// Port of the SFTP server.
	Port *int64 `json:"port,omitempty"`
	// Username to use to access the SFTP server.
	Username string `json:"username"`
}