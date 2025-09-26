package parse

import (
	"fmt"
	"testing"
)

func TestpeekNonSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		Root: &ListNode{
			Nodes: []Node{
				&TextNode{
					Text: []byte("test"),
				},
			},
		},
	}

	expected := item{
		typ: itemText,
		val: "test",
	}

	result := tree.peekNonSpace()

	if result.typ != expected.typ || result.val != expected.val {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
