package render

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPushValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		values: make(map[ast.NodeKind][]any),
	}

	ctx.PushValue(ast.NodeKind(1), "test")

	if len(ctx.values[ast.NodeKind(1)]) != 1 {
		t.Errorf("Expected length of values for NodeKind 1 to be 1, but got %d", len(ctx.values[ast.NodeKind(1)]))
	}

	if ctx.values[ast.NodeKind(1)][0] != "test" {
		t.Errorf("Expected value for NodeKind 1 to be 'test', but got %v", ctx.values[ast.NodeKind(1)][0])
	}
}
