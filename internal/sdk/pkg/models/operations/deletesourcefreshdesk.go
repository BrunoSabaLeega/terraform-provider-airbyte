// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceFreshdeskRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceFreshdeskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
