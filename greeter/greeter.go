package greeter

import (
	"fmt"
	"io"
)

func HelloTo(w io.Writer, name string) {
	msg := fmt.Sprintf("Hello, %s", name)
	_, _ = fmt.Fprintln(w, msg)
}
