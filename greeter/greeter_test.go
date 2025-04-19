package greeter_test

import (
	"bytes"
	"testing"

	"github.com/arth-go/mono/greeter"
)

func TestHelloPrintsTheMessage(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)

	greeter.HelloTo(buf, "World")
	want := "Hello, World\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
