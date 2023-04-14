// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum string

const (
	SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnumGoogleWorkspaceAdminReports SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum = "google-workspace-admin-reports"
)

func (e *SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "google-workspace-admin-reports":
		*e = SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum: %s", s)
	}
}

// SourceGoogleWorkspaceAdminReports - The values required to configure the source.
type SourceGoogleWorkspaceAdminReports struct {
	// The contents of the JSON service account key. See the <a href="https://developers.google.com/admin-sdk/reports/v1/guides/delegation">docs</a> for more information on how to generate this key.
	CredentialsJSON string `json:"credentials_json"`
	// The email of the user, which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
	// Sets the range of time shown in the report. Reports API allows from up to 180 days ago.
	Lookback   *int64                                                           `json:"lookback,omitempty"`
	SourceType SourceGoogleWorkspaceAdminReportsGoogleWorkspaceAdminReportsEnum `json:"sourceType"`
}