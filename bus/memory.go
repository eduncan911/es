package bus

import (
	"errors"
	"sync"
	"time"

	"github.com/eduncan911/es"
	"github.com/eduncan911/es/env"
)

type mem struct {
	c        chan es.Event
	handlers map[string]env.EventHandler
	started  bool

	mu sync.Mutex
}

// NewMem instantiates a new memory bus
func NewMem() Bus {
	return &mem{
		c:        make(chan es.Event, 10000),
		handlers: make(map[string]env.EventHandler),
	}
}

// AddEventHandler implements the Bus interface
func (m *mem) AddEventHandler(name string, h env.EventHandler) {
	m.handlers[name] = h
}

// Start implements the Bus interface
func (m *mem) Start() {
	if m.started {
		panic("Bus can't be started twice")
	}

	m.started = true

	go func() {
		for {
			select {
			case message := <-m.c:
				m.dispatch(message)
			}
		}
	}()
}

// MustPublish implements the env.Publisher contract
func (m *mem) MustPublish(es ...es.Event) {
	for _, e := range es {
		m.c <- e
	}
}

// Publish implements the env.Publisher contract
func (m *mem) Publish(es ...es.Event) error {
	m.MustPublish(es...)
	return nil
}

func (m *mem) dispatch(e es.Event) {
	for name, h := range m.handlers {
		if err := handleWithTimeout(h, e); err != nil {
			var contract = e.Meta().Contract
			l.Panicf("%s @ %s: %s", name, contract, err)
		}
	}
}

func handleWithTimeout(h env.EventHandler, e es.Event) (err error) {
	c := make(chan error, 1)

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	go func() {
		c <- h.HandleEvent(e)
	}()

	select {
	case err = <-c:
		return

	case <-time.After(time.Second):
		err = errors.New("Timeout")
	}

	return
}
