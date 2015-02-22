package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

// tracer is a Tracer that writes to an io.Writer.
type tracer struct {
	out io.Writer
}

// nilTracer
type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// New creates a new Tracer that will write the output to
// the specified io.Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off creates a tracer that will ignore all calls to Trace
func Off() Tracer {
	return &nilTracer{}
}
