package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEscapeAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{}
	c := context{
		state: stateAttr,
	}
	n := &parse.ActionNode{}
	c = e.escapeAction(c, n)
	if c.state != stateAttr {
		t.Errorf("expected stateAttr, got %v", c.state)
	}
}
