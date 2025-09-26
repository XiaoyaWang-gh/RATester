package parse

import (
	"fmt"
	"testing"
)

func TestcheckPipeline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
	}

	pipe := &PipeNode{
		Cmds: []*CommandNode{
			{
				Args: []Node{
					&StringNode{
						Text: "test",
					},
				},
			},
		},
	}

	tree.checkPipeline(pipe, "test")
}
