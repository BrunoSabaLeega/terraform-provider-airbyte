// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationSnowflakeRequest struct {
	DestinationSnowflakePutRequest *shared.DestinationSnowflakePutRequest `request:"mediaType=application/json"`
	DestinationID                  string                                 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationSnowflakeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}