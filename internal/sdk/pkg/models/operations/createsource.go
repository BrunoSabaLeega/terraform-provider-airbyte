// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/internal/sdk/pkg/models/shared"
)

type CreateSourceResponse struct {
	ContentType string
	// Successful operation
	SourceResponse *shared.SourceResponse
	StatusCode     int
	RawResponse    *http.Response
}