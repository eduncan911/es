package spec

import (
	"net/http"

	"github.com/eduncan911/es"
	"github.com/eduncan911/es/env"
)

type Values map[string]string

func GetJSON(url string, values Values) *env.Request {
	return &env.Request{
		Method:  "GET",
		Path:    url,
		Headers: nil,
		Body:    "",
	}
}
func PostJSON(url string, subj interface{}) *env.Request {
	return &env.Request{
		Method: "POST",
		Path:   url,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: subj,
	}
}

func PutJSON(url string, subj interface{}) *env.Request {
	return &env.Request{
		Method: "PUT",
		Path:   url,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: subj,
	}
}

func GivenEvents(events ...es.Event) []es.Event {
	if len(events) == 0 {
		return []es.Event{}
	}
	return events
}
