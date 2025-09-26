package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TesteditTemplateNode_1(t *testing.T) {
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
	callee := "testCallee"
	e.editTemplateNode(n, callee)
	if _, ok := e.templateNodeEdits[n]; !ok {
		t.Errorf("Expected node %s to be in templateNodeEdits", n)
	}
	if e.templateNodeEdits[n] != callee {
		t.Errorf("Expected callee %s, got %s", callee, e.templateNodeEdits[n])
	}
}
