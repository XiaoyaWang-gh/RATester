package web

import (
	"fmt"
	"testing"
)

func Testaddseg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := NewTree()
	tree.addseg([]string{"user", ":id"}, "user_id", []string{}, "")

	if len(tree.fixrouters) != 1 {
		t.Errorf("Expected 1 fixrouters, got %d", len(tree.fixrouters))
	}

	if tree.fixrouters[0].prefix != "user" {
		t.Errorf("Expected prefix 'user', got %s", tree.fixrouters[0].prefix)
	}

	if len(tree.fixrouters[0].leaves) != 1 {
		t.Errorf("Expected 1 leaves, got %d", len(tree.fixrouters[0].leaves))
	}

	if tree.fixrouters[0].leaves[0].runObject != "user_id" {
		t.Errorf("Expected runObject 'user_id', got %s", tree.fixrouters[0].leaves[0].runObject)
	}

	if len(tree.fixrouters[0].leaves[0].wildcards) != 1 {
		t.Errorf("Expected 1 wildcards, got %d", len(tree.fixrouters[0].leaves[0].wildcards))
	}

	if tree.fixrouters[0].leaves[0].wildcards[0] != ":id" {
		t.Errorf("Expected wildcard ':id', got %s", tree.fixrouters[0].leaves[0].wildcards[0])
	}
}
