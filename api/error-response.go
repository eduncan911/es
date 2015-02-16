package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code  int
	Error string
}

func (e *errorResponse) Write(w http.ResponseWriter) {

	type errDTO struct {
		Error string `json:"error"`
	}
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(&errDTO{e.Error})
	guard("marshal", err)
	w.WriteHeader(e.Code)
	w.Write(b)
}
