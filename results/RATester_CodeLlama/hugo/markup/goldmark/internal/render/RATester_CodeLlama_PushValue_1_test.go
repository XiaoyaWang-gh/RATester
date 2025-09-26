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

	ctx := &Context{}
	k := ast.NodeKind(1)
	v := 1
	ctx.PushValue(k, v)
	if ctx.values[k][0] != v {
		t.Errorf("ctx.values[k][0] = %v, want %v", ctx.values[k][0], v)
	}
}
