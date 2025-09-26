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

	kind := ast.NodeKind(1)
	expected := 0
	actual := ctx.GetAndIncrementOrdinal(kind)
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}

	expected = 1
	actual = ctx.GetAndIncrementOrdinal(kind)
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
