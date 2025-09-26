package http

import (
	"fmt"
	"testing"
)

func TestHeaderRegexp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := headerRegexp(tree, "key", "value")
	if err != nil {
		t.Fatal(err)
	}

	if !tree.match(nil) {
		t.Fatal("expected match")
	}

	if tree.matcher == nil {
		t.Fatal("expected matcher")
	}

	if tree.left != nil || tree.right != nil {
		t.Fatal("expected nil left and right")
	}

	if tree.operator != "" {
		t.Fatal("expected empty operator")
	}
}
