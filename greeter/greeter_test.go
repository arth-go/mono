package greeter_test

import (
	"bytes"
	"testing"

	"github.com/arth-go/mono/greeter"
)

func TestPrinterHelloWorld(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := greeter.NewPrinter()
	p.Output = buf

	p.Hello("World")

	want := "Hello, World\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
