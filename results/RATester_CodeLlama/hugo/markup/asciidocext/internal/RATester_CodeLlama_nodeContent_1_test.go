package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestNodeContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	node := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{Key: "class", Val: "foo"},
		},
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: "bar",
		},
	}
	want := "bar"
	got := nodeContent(node)
	if got != want {
		t.Errorf("nodeContent(%v) = %q; want %q", node, got, want)
	}
}
