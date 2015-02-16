package spec

import (
	"net/http"

	"github.com/eduncan911/es/env"
)

// ReturnErrorJSON defaults errors in JSON
func ReturnErrorJSON(status int, error string) *env.Response {
	return &env.Response{
		Status: status,
		Body: map[string]string{
			"error": error,
		},

		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

// ReturnJSON returns an `env.Response`
func ReturnJSON(response interface{}) *env.Response {
	return &env.Response{
		Status: http.StatusOK,
		Body:   response,

		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}
