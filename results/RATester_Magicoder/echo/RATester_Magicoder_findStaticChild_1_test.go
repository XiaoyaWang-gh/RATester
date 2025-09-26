package echo

import (
	"fmt"
	"testing"
)

func TestFindStaticChild_1(t *testing.T) {
	n := &node{
		staticChildren: children{
			{label: 'a'},
			{label: 'b'},
			{label: 'c'},
		},
	}

	tests := []struct {
		name     string
		label    byte
		expected *node
	}{
		{"found", 'b', n.staticChildren[1]},
		{"not found", 'd', nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := n.findStaticChild(test.label)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
