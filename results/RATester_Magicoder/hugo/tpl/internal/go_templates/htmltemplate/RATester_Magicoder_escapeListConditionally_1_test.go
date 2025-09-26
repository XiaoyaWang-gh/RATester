package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestescapeListConditionally_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		ns: &nameSpace{
			set: make(map[string]*Template),
		},
		output:            make(map[string]context),
		derived:           make(map[string]*template.Template),
		called:            make(map[string]bool),
		actionNodeEdits:   make(map[*parse.ActionNode][]string),
		templateNodeEdits: make(map[*parse.TemplateNode]string),
		textNodeEdits:     make(map[*parse.TextNode][]byte),
	}

	n := &parse.ListNode{
		Nodes: []parse.Node{
			&parse.ActionNode{},
			&parse.TemplateNode{},
			&parse.TextNode{},
		},
	}

	filter := func(e *escaper, c context) bool {
		return true
	}

	_, ok := e.escapeListConditionally(context{}, n, filter)

	if !ok {
		t.Errorf("Expected ok to be true, but got false")
	}
}
