package attributes

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	markdown := goldmark.New(
		goldmark.WithExtensions(&attrExtension{}),
	)

	source := []byte("Hello, World!")
	expected := "<p>Hello, World!</p>\n"

	var buf bytes.Buffer
	if err := markdown.Convert(source, &buf); err != nil {
		t.Fatal(err)
	}

	result := buf.String()
	if result != expected {
		t.Errorf("Expected: %q, but got: %q", expected, result)
	}
}
