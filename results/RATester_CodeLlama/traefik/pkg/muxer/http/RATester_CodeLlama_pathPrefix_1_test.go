package http

import (
	"fmt"
	"testing"
)

func TestPathPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := pathPrefix(tree, "/foo")
	if err != nil {
		t.Errorf("pathPrefix() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("pathPrefix() matcher = nil, want not nil")
	}

	if tree.operator != "" {
		t.Errorf("pathPrefix() operator = %q, want \"\"", tree.operator)
	}

	if tree.left != nil {
		t.Errorf("pathPrefix() left = %v, want nil", tree.left)
	}

	if tree.right != nil {
		t.Errorf("pathPrefix() right = %v, want nil", tree.right)
	}
}
