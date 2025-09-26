package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHTMLEscape_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w bytes.Buffer
	HTMLEscape(&w, []byte("Hello, <World>!"))
	if w.String() != "Hello, &lt;World&gt;!" {
		t.Errorf("HTMLEscape(%q) = %q, want %q", "Hello, <World>!", w.String(), "Hello, &lt;World&gt;!")
	}
}
