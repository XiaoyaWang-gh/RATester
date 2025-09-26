package tcp

import (
	"fmt"
	"testing"
)

func TestAlpn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	protos := []string{"h2"}
	err := alpn(tree, protos...)
	if err != nil {
		t.Errorf("alpn() failed: %v", err)
	}
	if tree.matcher == nil {
		t.Errorf("alpn() failed: matcher is nil")
	}
	if tree.operator != "||" {
		t.Errorf("alpn() failed: operator is %q, expected %q", tree.operator, "||")
	}
	if tree.left != nil {
		t.Errorf("alpn() failed: left is not nil")
	}
	if tree.right != nil {
		t.Errorf("alpn() failed: right is not nil")
	}
}
