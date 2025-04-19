package main

import (
	"github.com/arth-go/mono/greeter"
)

func main() {
	p := greeter.NewPrinter()

	p.Hello("World")
}
