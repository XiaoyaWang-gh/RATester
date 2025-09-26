package render

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPopValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		values: map[ast.NodeKind][]any{
			ast.KindText: {
				"foo",
				"bar",
			},
		},
	}
	v := ctx.PopValue(ast.KindText)
	if v != "bar" {
		t.Errorf("expected %q, but got %q", "bar", v)
	}
	if len(ctx.values[ast.KindText]) != 1 {
		t.Errorf("expected %d, but got %d", 1, len(ctx.values[ast.KindText]))
	}
	if ctx.values[ast.KindText][0] != "foo" {
		t.Errorf("expected %q, but got %q", "foo", ctx.values[ast.KindText][0])
	}
}
