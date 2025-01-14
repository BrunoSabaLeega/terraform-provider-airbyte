// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceKlaviyoRequest struct {
	SourceKlaviyoPutRequest *shared.SourceKlaviyoPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceKlaviyoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
