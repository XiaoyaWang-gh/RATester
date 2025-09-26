package goldmark

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuin/goldmark"
)

func TestExtend_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	markdown := goldmark.New(
		goldmark.WithExtensions(&tocExtension{}),
	)

	source := []byte("# Heading\n## Sub-heading")
	expected := []byte("# Heading\n## Sub-heading\n\n[TOC]\n\n## Sub-heading")

	var buf bytes.Buffer
	if err := markdown.Convert(source, &buf); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf.Bytes(), expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.Bytes())
	}
}
