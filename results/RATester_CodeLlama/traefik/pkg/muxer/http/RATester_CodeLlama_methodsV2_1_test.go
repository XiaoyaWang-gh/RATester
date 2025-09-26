package http

import (
	"fmt"
	"testing"
)

func TestMethodsV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := methodsV2(tree, "GET", "POST")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if tree.matcher == nil {
		t.Errorf("tree.matcher is nil")
	}
	if tree.operator != "OR" {
		t.Errorf("tree.operator is %s, want %s", tree.operator, "OR")
	}
	if tree.left == nil {
		t.Errorf("tree.left is nil")
	}
	if tree.right == nil {
		t.Errorf("tree.right is nil")
	}
}
