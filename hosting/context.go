package hosting

import (
	"github.com/eduncan911/es/api"
	"github.com/eduncan911/es/bus"
	"github.com/eduncan911/es/env"
	"github.com/gorilla/mux"
)

// Context represents a runtime's context of containers, modules, routes and event handlers
type Context struct {
	Items []*env.Container
}

// New instantiates a new Context by defining a new `env.Container` and registering
// that container with the module
func New(modules []env.Module) *Context {
	c := &Context{}
	for _, m := range modules {
		container := env.NewContainer()
		m.Register(container)
		c.Items = append(c.Items, container)
	}
	return c
}

// WireHTTP iterates over an instantiated context's `env.Container` collection to
// register each Route with the API's `mux.Router`
func (c *Context) WireHTTP(router *mux.Router) {
	for _, x := range c.Items {
		for _, route := range x.Routes {
			api.Handle(router, route)
		}
	}
}

// WireHandlers iterates over an instantiated context's `env.Container` collection to
// register each EventHandler on the application's defined `bus.Bus`
func (c *Context) WireHandlers(bus bus.Bus) {
	for _, x := range c.Items {
		for name, h := range x.Handlers {
			bus.AddEventHandler(name, h)
		}
	}
}

// func modules(modules []env.Module) *Context {
// 	var c = &Context{}
// 	for _, m := range modules {
// 		container := env.NewContainer()
// 		m.Register(container)
// 		c.Items = append(c.Items, container)
// 	}
// 	return c
// }
