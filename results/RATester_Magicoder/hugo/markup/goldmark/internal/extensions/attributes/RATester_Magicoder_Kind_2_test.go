package attributes

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestKind_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &attributesBlock{}
	expected := ast.NodeKind(kindAttributesBlock)
	if got := a.Kind(); got != expected {
		t.Errorf("Kind() = %v, want %v", got, expected)
	}
}
