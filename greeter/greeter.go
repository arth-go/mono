package greeter

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	output io.Writer
}
type Option func(*Printer)

func (p *Printer) Hello(name string) {
	_, _ = fmt.Fprintln(p.output, "Hello, "+name)
}
func WithOutput(output io.Writer) Option {
	return func(p *Printer) {
		p.output = output
	}
}

func NewPrinter(opts ...Option) *Printer {
	p := &Printer{
		output: os.Stdout,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
