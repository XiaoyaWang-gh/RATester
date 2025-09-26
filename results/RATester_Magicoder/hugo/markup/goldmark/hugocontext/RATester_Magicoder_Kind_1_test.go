package hugocontext

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestKind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoContext := &HugoContext{}
	expectedKind := ast.NodeKind(kindHugoContext)

	result := hugoContext.Kind()

	if result != expectedKind {
		t.Errorf("Expected Kind() to return %v, but got %v", expectedKind, result)
	}
}
