package parse

import (
	"fmt"
	"testing"
)

func TestIsEmptyTree_1(t *testing.T) {
	type testCase struct {
		name     string
		node     Node
		expected bool
	}

	testCases := []testCase{
		{
			name:     "NilNode",
			node:     nil,
			expected: true,
		},
		{
			name:     "ActionNode",
			node:     &ActionNode{},
			expected: false,
		},
		{
			name:     "CommentNode",
			node:     &CommentNode{},
			expected: true,
		},
		{
			name:     "IfNode",
			node:     &IfNode{},
			expected: false,
		},
		{
			name:     "ListNode",
			node:     &ListNode{},
			expected: true,
		},
		{
			name:     "RangeNode",
			node:     &RangeNode{},
			expected: false,
		},
		{
			name:     "TemplateNode",
			node:     &TemplateNode{},
			expected: false,
		},
		{
			name:     "TextNode",
			node:     &TextNode{},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := IsEmptyTree(tc.node)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
