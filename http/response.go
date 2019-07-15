package http

import (
	"net/http"
)

// Response is an HTTP response object
type Response http.Response

// Native downcasts the response to the native HTTP interface
func (r *Response) Native() *http.Response {
	return (*http.Response)(r)
}

// NewResponse creates a new HTTP response object
func NewResponse(response *http.Response) *Response {
	return (*Response)(response)
}
