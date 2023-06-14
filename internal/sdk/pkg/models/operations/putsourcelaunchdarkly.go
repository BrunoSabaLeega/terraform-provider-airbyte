// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceLaunchdarklyRequest struct {
	SourceLaunchdarklyPutRequest *shared.SourceLaunchdarklyPutRequest `request:"mediaType=application/json"`
	SourceID                     string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceLaunchdarklyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}