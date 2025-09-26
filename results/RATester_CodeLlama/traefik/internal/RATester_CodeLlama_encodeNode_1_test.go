package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/paerser/parser"
)

func TestEncodeNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	labels := make(map[string]string)
	root := "root"
	node := &parser.Node{
		Name: "node",
		Children: []*parser.Node{
			{
				Name: "child1",
				Children: []*parser.Node{
					{
						Name: "grandchild1",
					},
				},
			},
			{
				Name: "child2",
				Children: []*parser.Node{
					{
						Name: "grandchild2",
					},
				},
			},
		},
	}

	// when
	encodeNode(labels, root, node)

	// then
	assert.Equal(t, map[string]string{
		"root.node.child1.grandchild1": "",
		"root.node.child2.grandchild2": "",
	}, labels)
}
