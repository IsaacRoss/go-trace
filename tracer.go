package trace

import (
	"fmt"
	"io"
)

//Tracer is the interface that describes an object capable of tracing
// events trhoughout code

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}
type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
