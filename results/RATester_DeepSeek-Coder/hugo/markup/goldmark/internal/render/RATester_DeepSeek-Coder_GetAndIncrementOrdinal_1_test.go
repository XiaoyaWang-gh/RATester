package render

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestGetAndIncrementOrdinal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		ordinals: make(map[ast.NodeKind]int),
	}

	kind1 := ast.NodeKind(1)
	kind2 := ast.NodeKind(2)

	expected1 := 0
	expected2 := 0

	result1 := ctx.GetAndIncrementOrdinal(kind1)
	result2 := ctx.GetAndIncrementOrdinal(kind2)
	result3 := ctx.GetAndIncrementOrdinal(kind1)

	if result1 != expected1 {
		t.Errorf("Expected %d, got %d", expected1, result1)
	}
	if result2 != expected2 {
		t.Errorf("Expected %d, got %d", expected2, result2)
	}
	if result3 != 1 {
		t.Errorf("Expected 1, got %d", result3)
	}
}
