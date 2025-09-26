package doctree

import (
	"fmt"
	"testing"
)

func TestData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &WalkContext[any]{}
	tree := ctx.Data()
	if tree == nil {
		t.Error("Expected non-nil tree, got nil")
	}
}
