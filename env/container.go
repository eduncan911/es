package env

import (
	"github.com/eduncan911/es/api"
)

// Container represents the code contract implementation for a module
type Container struct {
	Routes    []*api.Route
	Handlers  EventHandlerMap
	DataReset map[string]func()
}

// NewContainer instantiates a new container
func NewContainer() *Container {
	return &Container{
		Routes:    []*api.Route{},
		Handlers:  EventHandlerMap{},
		DataReset: make(map[string]func()),
	}
}

// HandleHTTP registers http handlers with api.Routes internally
func (c *Container) HandleHTTP(
	method string,
	path string,
	handler api.Handler,
) {
	c.Routes = append(c.Routes, api.NewRoute(method, path, handler))
}

// HandleEvents registers the event handler internally
func (c *Container) HandleEvents(
	name string,
	handler EventHandler,
) {
	c.Handlers[name] = handler
}

// ResetData is used to register func()s to rollback data changes
func (c *Container) ResetData(
	name string,
	reset func(),
) {
	c.DataReset[name] = reset
}
