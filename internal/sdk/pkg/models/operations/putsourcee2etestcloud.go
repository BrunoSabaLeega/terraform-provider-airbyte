// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceE2eTestCloudRequest struct {
	SourceE2eTestCloudPutRequest *shared.SourceE2eTestCloudPutRequest `request:"mediaType=application/json"`
	SourceID                     string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceE2eTestCloudResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}