package parse

import (
	"fmt"
	"testing"
)

func TestNewIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ident := "test"
	node := NewIdentifier(ident)

	if node.Ident != ident {
		t.Errorf("Expected Ident to be %s, but got %s", ident, node.Ident)
	}

	if node.NodeType != NodeIdentifier {
		t.Errorf("Expected NodeType to be %d, but got %d", NodeIdentifier, node.NodeType)
	}
}
