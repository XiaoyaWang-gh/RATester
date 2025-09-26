package tplimpl

import (
	"context"
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/htmltemplate"
)

func TestInit_34(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	tmpl := &template.Template{}
	var tt templateExecHelper
	tt.Init(ctx, tmpl)
}
