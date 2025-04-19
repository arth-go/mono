package greeter

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Hello(name string) {
	_, _ = fmt.Fprintln(p.Output, "Hello, "+name)
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}
