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

	tree := &Tree{}
	node := &IdentifierNode{}
	node.SetTree(tree)

	if node.tr != tree {
		t.Errorf("Expected tree to be set, but it was not.")
	}
}
