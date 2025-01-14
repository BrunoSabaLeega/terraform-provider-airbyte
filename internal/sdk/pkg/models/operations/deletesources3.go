// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceS3Request struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceS3Response struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
