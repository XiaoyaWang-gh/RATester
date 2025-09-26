package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEditTextNode_1(t *testing.T) {
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

	e.editTextNode(n, []byte("edited text"))

	if _, ok := e.textNodeEdits[n]; !ok {
		t.Errorf("Expected textNodeEdits to contain node %v", n)
	}

	if string(e.textNodeEdits[n]) != "edited text" {
		t.Errorf("Expected edited text to be 'edited text', got %s", string(e.textNodeEdits[n]))
	}
}
