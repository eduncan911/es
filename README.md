
# es
    import "github.com/eduncan911/es"

Package es is an event-sourcing infrastructure and specs framework for GoLang originally created by Rinat Abdullin (a href="https://github.com/abdullin/omni/">https://github.com/abdullin/omni/</a>) and modified by Eric Duncan (a href="https://github.com/eduncan911/es/">https://github.com/eduncan911/es/</a>).

This folder contains the core infrastructure for prototyping event-driven back-ends. You can import it in your go and move from there.

Additional packages available are:


	* `root` - binary-sortable UUID and a definition of an event
	* `api` - logic for hosting a simple JSON API (with some helpers)
	* `bus` - event bus and an in-memory implementation
	* `log` - helpers to setup logging
	* `env` - environment for defining modules and specs (contracts)
	* `specs` - express, verify and print event-driven specifications
	* `hosting` - wire and run modules in a process







## type Contract
``` go
type Contract string
```
Contract is the string name











## type Event
``` go
type Event interface {
    Meta() *Info
}
```
Event interface defines the Meta info











## type Info
``` go
type Info struct {
    Contract string
    EventID  uuid.ID
}
```
Info represents a Contract and EventID









### func NewInfo
``` go
func NewInfo(contract string, eventID uuid.ID) *Info
```
NewInfo instantiates a new Info










- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)