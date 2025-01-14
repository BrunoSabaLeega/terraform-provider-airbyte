// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteWorkspaceRequest struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceId"`
}

type DeleteWorkspaceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
