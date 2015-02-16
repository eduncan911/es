package es

import (
	"github.com/eduncan911/go/pkg/uuid"
)

// Contract is the string name
type Contract string

// Event interface defines the Meta info
type Event interface {
	Meta() *Info
}

// Info represents a Contract and EventID
type Info struct {
	Contract string
	EventID  uuid.ID
}

// NewInfo instantiates a new Info
func NewInfo(contract string, eventID uuid.ID) *Info {
	return &Info{
		Contract: contract,
		EventID:  eventID,
	}
}
