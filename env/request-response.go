package env

import (
	"net/http"
)

// Request represents the method, path, headers and body of the request
type Request struct {
	Method  string
	Path    string
	Headers http.Header
	Body    interface{}
}

// Response represents an http response that encapsulates the status, headers and body
type Response struct {
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Body    interface{} `json:"body"`
}
