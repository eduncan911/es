package spec

import (
	"github.com/eduncan911/es"
)

func newPublisher() *publisher {
	return &publisher{}
}

type publisher struct {
	Events []es.Event
}

func (p *publisher) MustPublish(es ...es.Event) {
	if err := p.Publish(es...); err != nil {
		panic(err)
	}
}

func (p *publisher) Publish(es ...es.Event) error {
	p.Events = append(p.Events, es...)
	return nil
}

func (p *publisher) Clear() {
	p.Events = nil
}
