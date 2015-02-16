package spec

import "github.com/eduncan911/es"

// Events converts one or more `es.Event` to an array of `es.Event`s
func Events(events ...es.Event) []es.Event {
	if len(events) == 0 {
		return []es.Event{}
	}
	return events

}
