package bus

import (
	"github.com/eduncan911/es/env"
)

// Bus interface defines the service bus
type Bus interface {
	// AddEventHandler takes the name and handler
	AddEventHandler(name string, h env.EventHandler)
	// Publisher enforces the publishing requirements for the bus
	env.Publisher
	// Start initializes the bus' internals
	Start()
}
