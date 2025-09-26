package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEditTemplateNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		templateNodeEdits: make(map[*parse.TemplateNode]string),
	}

	n := &parse.TemplateNode{
		Name: "test",
	}

	e.editTemplateNode(n, "callee")

	if _, ok := e.templateNodeEdits[n]; !ok {
		t.Errorf("Expected node to be in templateNodeEdits map")
	}

	if e.templateNodeEdits[n] != "callee" {
		t.Errorf("Expected callee to be stored in templateNodeEdits map")
	}
}
