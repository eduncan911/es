/*Package hosting defines a `hosting.New(...)` helper method for use in creating the runtime by taking in modules and registering their API routes and Event Handlers.

Example:

	package main

	import (
		"log"
		"net/http"

		"github.com/eduncan911/es/hosting"
		"github.com/eduncan911/es/bus"
		"github.com/eduncan911/es/env"

		"github.com/gorilla/mux"
	)

	func main() {

		router := mux.NewRouter()
		memBus := bus.NewMem()

		var wrap = bus.WrapWithLogging(memBus)
		modules := factory(wrap)

		// call the hosting.New(...) helper method
		//
		var host = hosting.New(modules)

		host.WireHttp(router)
		host.WireHandlers(memBus)
		memBus.Start()

		bind := ":8001"
		log.Info("Listening at %v", bind)
		log.Panic(http.ListenAndServe(bind, router))
	}

	func factory(pub env.Publisher) []env.Module {
		return []env.Module{
			views.NewModule(pub),
			task.NewModule(pub),
		}
	}


*/
package hosting
