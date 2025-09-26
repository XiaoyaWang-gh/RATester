package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestAttr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	node := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{
				Key: "id",
				Val: "my-id",
			},
			{
				Key: "class",
				Val: "my-class",
			},
		},
	}
	key := "id"
	want := "my-id"
	got := attr(node, key)
	if got != want {
		t.Errorf("attr(%v, %v) = %v; want %v", node, key, got, want)
	}
}
