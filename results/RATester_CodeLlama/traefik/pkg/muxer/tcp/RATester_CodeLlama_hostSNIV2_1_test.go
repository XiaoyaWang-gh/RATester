package tcp

import (
	"fmt"
	"testing"
)

func TestHostSNIV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := hostSNIV2(tree, "example.com")
	if err != nil {
		t.Error(err)
	}
	if tree.matcher == nil {
		t.Error("matcher should not be nil")
	}
	if tree.left != nil || tree.right != nil {
		t.Error("left and right should be nil")
	}
	if tree.operator != "" {
		t.Error("operator should be empty")
	}
}
