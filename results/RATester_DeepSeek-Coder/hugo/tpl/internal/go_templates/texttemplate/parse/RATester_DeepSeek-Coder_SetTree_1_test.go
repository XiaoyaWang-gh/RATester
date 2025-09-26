package parse

import (
	"fmt"
	"testing"
)

func TestSetTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{Name: "test"}
	identifier := &IdentifierNode{Ident: "testIdent"}

	identifier.SetTree(tree)

	if identifier.tr != tree {
		t.Errorf("Expected tree to be %v, got %v", tree, identifier.tr)
	}
}
