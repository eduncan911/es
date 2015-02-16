package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Response interface defines the Write() method for the ResponseWriter
type Response interface {
	Write(w http.ResponseWriter)
}

// Request represents an http Request
type Request struct {
	Raw *http.Request
}

// NewRequest instantiates a new Request
func NewRequest(inner *http.Request) *Request {
	return &Request{inner}
}

// String returns the value of the request's FormValue, or panics if not available
func (r *Request) String(param string) string {
	// TODO Panic?
	if v := r.Raw.FormValue(param); v == "" {
		panic(NewErrArgumentNil(param))
	} else {
		return v
	}
}

// ParseBody will encode the body for Content-Type == json.  Others can be added here, such as XML.
func (r *Request) ParseBody(subj interface{}) error {
	content := r.Raw.Header.Get("Content-Type")
	switch content {
	case "application/json":
		decoder := json.NewDecoder(r.Raw.Body)
		return decoder.Decode(subj)
	}

	return errors.New("Unexpected content type " + content)
}
