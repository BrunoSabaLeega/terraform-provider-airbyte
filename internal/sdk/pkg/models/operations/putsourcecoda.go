// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceCodaRequest struct {
	SourceCodaPutRequest *shared.SourceCodaPutRequest `request:"mediaType=application/json"`
	SourceID             string                       `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceCodaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
