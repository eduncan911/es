package bus

import (
	"fmt"

	"github.com/eduncan911/es"
	"github.com/eduncan911/es/env"
	"github.com/op/go-logging"
)

var (
	l = logging.MustGetLogger("bus")
)

// LoggingWrapper represents an env.Publisher with logging
type LoggingWrapper struct {
	inner env.Publisher
}

// WrapWithLogging instantiates a new LoggingWrapper with the specified env.Publisher
func WrapWithLogging(p env.Publisher) env.Publisher {
	return &LoggingWrapper{p}
}

// Publish takes one or more events, logs the event, and publishes them to the env.Publisher that returns an error if one occurs
func (p *LoggingWrapper) Publish(e ...es.Event) error {
	log(e...)
	return p.inner.Publish(e...)
}

// MustPublish takes one or more events, logs the event, and publishes them to the env.Publisher
func (p *LoggingWrapper) MustPublish(e ...es.Event) {
	log(e...)
	p.inner.MustPublish(e...)
}

func log(es ...es.Event) {
	for _, e := range es {
		switch e := e.(type) {
		case fmt.Stringer:
			l.Debug("%v", e.String())
		default:
			var contract = e.Meta().Contract
			l.Debug("%v", contract)
		}
	}
}
