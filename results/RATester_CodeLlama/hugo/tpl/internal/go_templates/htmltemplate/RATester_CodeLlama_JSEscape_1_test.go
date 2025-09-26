package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestJSEscape_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w bytes.Buffer
	JSEscape(&w, []byte("hello"))
	if w.String() != "hello" {
		t.Errorf("JSEscape(%q) = %q, want %q", "hello", w.String(), "hello")
	}
}
