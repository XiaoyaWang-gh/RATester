package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestescapeText_1(t *testing.T) {
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

	n := &parse.TextNode{
		Text: []byte("<test>"),
	}

	ctx := context{
		state: stateText,
	}

	ctx = e.escapeText(ctx, n)

	if ctx.state != stateText {
		t.Errorf("Expected stateText, got %v", ctx.state)
	}

	if string(n.Text) != "&lt;test&gt;" {
		t.Errorf("Expected escaped text, got %s", string(n.Text))
	}
}
