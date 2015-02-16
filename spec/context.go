package spec

import (
	"github.com/eduncan911/es/env"
)

// Context represents the spec's definition
type Context struct {
	pub  *publisher
	spec *env.Spec
}

// NewContext instantiates a new Context with the specified `env.Spec`
func NewContext(spec *env.Spec) *Context {
	return &Context{
		pub:  newPublisher(),
		spec: spec,
	}
}

// Pub represents the `env.Publisher`
func (c *Context) Pub() env.Publisher {
	return c.pub
}

// Verify returns a `Report` on on the specified `end.Module`
func (c *Context) Verify(m env.Module) *Report {
	return buildAndVerify(c.pub, c.spec, m)
}
