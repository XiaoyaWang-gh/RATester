package parse

import (
	"fmt"
	"testing"
)

func TestIsEmptyTree_1(t *testing.T) {
	tests := []struct {
		name string
		n    Node
		want bool
	}{
		{
			name: "nil",
			n:    nil,
			want: true,
		},
		{
			name: "ActionNode",
			n:    &ActionNode{},
			want: false,
		},
		{
			name: "CommentNode",
			n:    &CommentNode{},
			want: true,
		},
		{
			name: "IfNode",
			n:    &IfNode{},
			want: false,
		},
		{
			name: "ListNode",
			n:    &ListNode{},
			want: false,
		},
		{
			name: "RangeNode",
			n:    &RangeNode{},
			want: false,
		},
		{
			name: "TemplateNode",
			n:    &TemplateNode{},
			want: false,
		},
		{
			name: "TextNode",
			n:    &TextNode{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsEmptyTree(tt.n); got != tt.want {
				t.Errorf("IsEmptyTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
