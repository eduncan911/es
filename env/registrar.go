package env

import (
	"github.com/eduncan911/es/api"
)

// Registrar interface implements registration for a module
type Registrar interface {
	HandleHttp(method string, path string, handler api.Handler)
	HandleEvents(name string, handler EventHandler)
	ResetData(name string, action func())
}
