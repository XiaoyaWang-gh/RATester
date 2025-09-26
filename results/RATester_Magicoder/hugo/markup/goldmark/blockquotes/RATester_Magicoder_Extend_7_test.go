package blockquotes

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	markdown := goldmark.New(
		goldmark.WithExtensions(&blockquotesExtension{}),
	)

	source := []byte("> This is a blockquote.")
	expected := "<blockquote>\n<p>This is a blockquote.</p>\n</blockquote>\n"

	var buf bytes.Buffer
	if err := markdown.Convert(source, &buf); err != nil {
		t.Fatal(err)
	}

	if buf.String() != expected {
		t.Errorf("Unexpected output from Markdown.Convert:\nExpected: %q\nGot: %q", expected, buf.String())
	}
}
