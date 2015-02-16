package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handle defines a function to wrap a route a route
func Handle(router *mux.Router, route *Route) *mux.Route {

	// capture panics and don't let the web app crash
	var handler = func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					handleError(w, err)
					return
				}
				panic(r)
			}
		}()
	}

	return router.Methods(route.Method).Subrouter().HandleFunc(route.Path, handler)
}
