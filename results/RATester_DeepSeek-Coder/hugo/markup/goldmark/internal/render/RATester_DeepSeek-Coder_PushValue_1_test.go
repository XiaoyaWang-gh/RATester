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

	testKind := ast.NodeKind(1)
	testValue := "test"

	ctx.PushValue(testKind, testValue)

	if len(ctx.values[testKind]) != 1 {
		t.Errorf("Expected length of values for kind %v to be 1, got %v", testKind, len(ctx.values[testKind]))
	}

	if ctx.values[testKind][0] != testValue {
		t.Errorf("Expected value at index 0 of values for kind %v to be %v, got %v", testKind, testValue, ctx.values[testKind][0])
	}
}
