package web

import (
	"fmt"
	"testing"
)

func TestfindAndRemoveSingleTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		prefix: "/",
		fixrouters: []*Tree{
			{
				prefix: "/test",
				leaves: []*leafInfo{
					{
						runObject: "test",
					},
				},
			},
		},
	}

	findAndRemoveSingleTree(tree)

	if len(tree.fixrouters) != 0 {
		t.Errorf("Expected fixrouters to be empty, but got %d", len(tree.fixrouters))
	}

	if len(tree.leaves) != 0 {
		t.Errorf("Expected leaves to be empty, but got %d", len(tree.leaves))
	}
}
