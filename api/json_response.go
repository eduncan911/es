package api

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Code int

	subj interface{}
}

// NewJSON instantiates a new JSON encoded response
func NewJSON(subj interface{}) Response {
	return &jsonResponse{
		Code: 200,
		subj: subj,
	}
}

// Write implements the Response interface
func (e *jsonResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(e.subj)
	guard("marshal", err)
	w.WriteHeader(e.Code)
	w.Write(b)
}
