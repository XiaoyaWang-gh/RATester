package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOperand_1(t *testing.T) {
	type testCase struct {
		name     string
		setup    func(t *Tree)
		expected Node
	}

	tests := []testCase{
		{
			name: "test case 1",
			setup: func(t *Tree) {
				// setup code here
			},
			expected: nil, // expected result
		},
		// more test cases...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tree := &Tree{}
			if tc.setup != nil {
				tc.setup(tree)
			}
			result := tree.operand()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
