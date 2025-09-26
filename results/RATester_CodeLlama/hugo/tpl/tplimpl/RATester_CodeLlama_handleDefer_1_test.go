package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestHandleDefer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	withNode := &parse.WithNode{}
	c := &templateContext{}
	c.handleDefer(withNode)
}
