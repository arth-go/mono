package main

import (
	"os"

	"github.com/arth-go/mono/greeter"
)

func main() {
	greeter.HelloTo(os.Stdout, "World")
}
