package web

import (
	"fmt"
	"testing"
)

func Testaddtree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	segments := []string{"test", ":param"}
	wildcards := []string{}
	reg := ""
	tree.addtree(segments, tree, wildcards, reg)

	if tree.prefix != "test" {
		t.Errorf("Expected prefix to be 'test', but got %s", tree.prefix)
	}

	if len(tree.fixrouters) != 1 {
		t.Errorf("Expected 1 fixrouters, but got %d", len(tree.fixrouters))
	}

	if tree.wildcard != nil {
		t.Errorf("Expected wildcard to be nil, but got %v", tree.wildcard)
	}

	if len(tree.leaves) != 0 {
		t.Errorf("Expected 0 leaves, but got %d", len(tree.leaves))
	}
}
