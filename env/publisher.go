package env

import (
	"github.com/eduncan911/es"
)

// Publisher is an interface for publishing contracts
type Publisher interface {
	// Publish pushes one or more events onto the bus
	Publish(e ...es.Event) error
	// MustPublish pushes one or more events onto the bus with internal error handling
	MustPublish(e ...es.Event)
}
