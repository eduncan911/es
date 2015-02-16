
# api
    import "github.com/eduncan911/es/api"

Package api defines the API code contracts for registering and serving requests in an API.






## func Handle
``` go
func Handle(router *mux.Router, route *Route) *mux.Route
```
Handle defines a function to wrap a route a route



## type ErrArgument
``` go
type ErrArgument struct {
    Argument string
    Problem  string
}
```
ErrArgument represents a single specific error in an error collection









### func NewErrArgument
``` go
func NewErrArgument(param, problem string) *ErrArgument
```
NewErrArgument instantiates a new ErrArugment with the specified problem


### func NewErrArgumentNil
``` go
func NewErrArgumentNil(param string) *ErrArgument
```
NewErrArgumentNil instantiates a new ErrArgument for a missing parameter




### func (\*ErrArgument) Error
``` go
func (e *ErrArgument) Error() string
```
Error returns the string representation of the ErrArgument



## type Handler
``` go
type Handler func(r *Request) Response
```
Handler defines the function that implements the request to response











## type Request
``` go
type Request struct {
    Raw *http.Request
}
```
Request represents an http Request









### func NewRequest
``` go
func NewRequest(inner *http.Request) *Request
```
NewRequest instantiates a new Request




### func (\*Request) ParseBody
``` go
func (r *Request) ParseBody(subj interface{}) error
```
ParseBody will encode the body for Content-Type == json.  Others can be added here, such as XML.



### func (\*Request) String
``` go
func (r *Request) String(param string) string
```
String returns the value of the request's FormValue, or panics if not available



## type Response
``` go
type Response interface {
    Write(w http.ResponseWriter)
}
```
Response interface defines the Write() method for the ResponseWriter









### func BadRequestResponse
``` go
func BadRequestResponse(err string) Response
```
BadRequestResponse instantiates a new bad request Response


### func NewError
``` go
func NewError(err string, code int) Response
```
NewError instantiates a new error Response


### func NewJSON
``` go
func NewJSON(subj interface{}) Response
```
NewJSON instantiates a new JSON encoded response


### func NotImplementedResponse
``` go
func NotImplementedResponse() Response
```
NotImplementedResponse instantiates a new not implemented Response




## type Route
``` go
type Route struct {
    Method  string
    Path    string
    Handler Handler
}
```
Route represents an API route's method, path and Handler func









### func NewRoute
``` go
func NewRoute(method, path string, handler Handler) *Route
```
NewRoute returns a new instance of a Route with the defined parameters










- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)