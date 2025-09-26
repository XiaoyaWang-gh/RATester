package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEscapeList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{}
	c := context{}
	n := &parse.ListNode{}
	c = e.escapeList(c, n)
	if c.state != stateDead {
		t.Errorf("expected stateDead, got %v", c.state)
	}
}
