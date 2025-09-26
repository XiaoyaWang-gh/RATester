package http

import (
	"fmt"
	"testing"
)

func TestPathV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	paths := []string{"/path1", "/path2"}

	if err := pathV2(tree, paths...); err != nil {
		t.Errorf("pathV2() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("pathV2() matcher = nil, want not nil")
	}

	if tree.operator != "OR" {
		t.Errorf("pathV2() operator = %s, want %s", tree.operator, "OR")
	}

	if tree.left == nil {
		t.Errorf("pathV2() left = nil, want not nil")
	}

	if tree.right == nil {
		t.Errorf("pathV2() right = nil, want not nil")
	}
}
