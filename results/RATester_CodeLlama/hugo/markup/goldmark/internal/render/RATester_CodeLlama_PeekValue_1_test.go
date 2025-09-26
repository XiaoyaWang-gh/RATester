package render

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPeekValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{}
	ctx.values = map[ast.NodeKind][]any{
		ast.KindText: []any{"foo", "bar"},
	}
	if got := ctx.PeekValue(ast.KindText); got != "bar" {
		t.Errorf("PeekValue() = %v, want %v", got, "bar")
	}
}
