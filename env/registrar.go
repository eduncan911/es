package env

import (
	"github.com/eduncan911/es/api"
)

// Registrar interface implements registration for a module
type Registrar interface {
	// HandleHTTP registers http handlers with api.Routes internally
	HandleHTTP(method string, path string, handler api.Handler)
	// HandleEvents registers the event handler internally
	HandleEvents(name string, handler EventHandler)
	// ResetData is used to register func()s to rollback data changes
	ResetData(name string, action func())
}
