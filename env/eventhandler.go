package env

import (
	"github.com/eduncan911/es"
)

// EventHandler interface represents an event handler
type EventHandler interface {
	// Contracts returns a string array of contracts for this handler
	Contracts() []string
	// HandlerEvent handles the event
	HandleEvent(e es.Event) error
}

// EventHandlerMap is a map of regstered event handlers with a string as the key
type EventHandlerMap map[string]EventHandler
