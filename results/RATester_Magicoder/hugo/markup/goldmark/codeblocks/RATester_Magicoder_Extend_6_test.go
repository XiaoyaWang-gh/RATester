package codeblocks

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	markdown := goldmark.New(
		goldmark.WithExtensions(&codeBlocksExtension{}),
	)

	source := []byte("```go\nfunc main() {\n\tprintln(\"Hello, World!\") \n}\n```")
	expected := "<pre><code class=\"language-go\">func main() {\n\tprintln(\"Hello, World!\") \n}\n</code></pre>"

	var buf bytes.Buffer
	if err := markdown.Convert(source, &buf); err != nil {
		t.Fatal(err)
	}

	result := buf.String()
	if result != expected {
		t.Errorf("Expected: %q, but got: %q", expected, result)
	}
}
