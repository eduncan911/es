package api

import (
	"fmt"
	"net/http"
)

// NewError instantiates a new error Response
func NewError(err string, code int) Response {
	return &errorResponse{
		Code:  code,
		Error: err,
	}
}

// BadRequestResponse instantiates a new bad request Response
func BadRequestResponse(err string) Response {
	return &errorResponse{
		Code:  http.StatusBadRequest,
		Error: err,
	}
}

// NotImplementedResponse instantiates a new not implemented Response
func NotImplementedResponse() Response {
	return &errorResponse{
		Code:  http.StatusNotImplemented,
		Error: "Not implemented",
	}
}

func handleError(w http.ResponseWriter, err error) {
	switch err := err.(type) {
	case *ErrArgument:
		NewError(err.Error(), http.StatusBadRequest).Write(w)
	}
}

// ErrArgument represents a single specific error in an error collection
type ErrArgument struct {
	Argument string
	Problem  string
}

// Error returns the string representation of the ErrArgument
func (e *ErrArgument) Error() string {
	return fmt.Sprintf("Argument '%v': %v", e.Argument, e.Problem)
}

// NewErrArgument instantiates a new ErrArugment with the specified problem
func NewErrArgument(param, problem string) *ErrArgument {
	return &ErrArgument{
		Argument: param,
		Problem:  problem,
	}
}

// NewErrArgumentNil instantiates a new ErrArgument for a missing parameter
func NewErrArgumentNil(param string) *ErrArgument {
	return &ErrArgument{
		Argument: param,
		Problem:  "can't be empty",
	}
}
