package api

// Route represents an API route's method, path and Handler func
type Route struct {
	Method  string
	Path    string
	Handler Handler
}

// Handler defines the function that implements the request to response
type Handler func(r *Request) Response

// NewRoute returns a new instance of a Route with the defined parameters
func NewRoute(method, path string, handler Handler) *Route {
	return &Route{method, path, handler}
}
