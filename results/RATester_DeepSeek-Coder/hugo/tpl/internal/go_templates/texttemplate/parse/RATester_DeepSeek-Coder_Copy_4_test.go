package parse

import (
	"fmt"
	"testing"
)

func TestCopy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name:      "test",
		ParseName: "test",
		Root: &ListNode{
			Nodes: []Node{
				&StringNode{
					Quoted: "\"test\"",
					Text:   "test",
				},
			},
		},
	}

	copyTree := tree.Copy()

	if copyTree.Name != tree.Name {
		t.Errorf("Expected Name to be %s, got %s", tree.Name, copyTree.Name)
	}

	if copyTree.ParseName != tree.ParseName {
		t.Errorf("Expected ParseName to be %s, got %s", tree.ParseName, copyTree.ParseName)
	}

	if copyTree.Root.Nodes[0].(*StringNode).Quoted != tree.Root.Nodes[0].(*StringNode).Quoted {
		t.Errorf("Expected Quoted to be %s, got %s", tree.Root.Nodes[0].(*StringNode).Quoted, copyTree.Root.Nodes[0].(*StringNode).Quoted)
	}

	if copyTree.Root.Nodes[0].(*StringNode).Text != tree.Root.Nodes[0].(*StringNode).Text {
		t.Errorf("Expected Text to be %s, got %s", tree.Root.Nodes[0].(*StringNode).Text, copyTree.Root.Nodes[0].(*StringNode).Text)
	}
}
