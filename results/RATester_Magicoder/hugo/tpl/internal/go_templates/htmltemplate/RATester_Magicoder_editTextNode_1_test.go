package template

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TesteditTextNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		textNodeEdits: make(map[*parse.TextNode][]byte),
	}
	n := &parse.TextNode{
		Text: []byte("test text"),
	}
	text := []byte("edited text")
	e.editTextNode(n, text)
	if _, ok := e.textNodeEdits[n]; !ok {
		t.Errorf("Expected textNodeEdits to contain node %v", n)
	}
	if !bytes.Equal(e.textNodeEdits[n], text) {
		t.Errorf("Expected textNodeEdits to contain edited text %s, got %s", text, e.textNodeEdits[n])
	}
}
