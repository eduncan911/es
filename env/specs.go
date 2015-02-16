package env

import (
	"github.com/eduncan911/es"
)

// Spec implements a specification for the UseCase
type Spec struct {
	Name string
	//Factory Factory
	UseCases []UseCaseFactory
}

// UseCaseFactory is the function to execute the UseCase
type UseCaseFactory func() *UseCase

// UseCase represents a BDD specification
type UseCase struct {
	Name string

	Given        []es.Event
	When         *Request
	ThenEvents   []es.Event
	ThenResponse *Response

	Where Where
}

// Where defines a map of conditions
type Where interface {
	Map() map[string]string
}
