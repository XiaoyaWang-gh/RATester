package http

import (
	"fmt"
	"testing"
)

func TestPathPrefixV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	paths := []string{"/path1", "/path2"}

	err := pathPrefixV2(tree, paths...)
	if err != nil {
		t.Errorf("pathPrefixV2() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("pathPrefixV2() matcher = nil, want not nil")
	}

	if tree.operator != "OR" {
		t.Errorf("pathPrefixV2() operator = %s, want %s", tree.operator, "OR")
	}

	if tree.left == nil {
		t.Errorf("pathPrefixV2() left = nil, want not nil")
	}

	if tree.right == nil {
		t.Errorf("pathPrefixV2() right = nil, want not nil")
	}
}
