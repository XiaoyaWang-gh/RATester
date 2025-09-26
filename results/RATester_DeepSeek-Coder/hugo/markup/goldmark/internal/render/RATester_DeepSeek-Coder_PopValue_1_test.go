package render

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPopValue_1(t *testing.T) {
	type testCase struct {
		name     string
		ctx      *Context
		kind     ast.NodeKind
		expected any
	}

	testCases := []testCase{
		{
			name: "PopValue from non-empty slice",
			ctx: &Context{
				values: map[ast.NodeKind][]any{
					ast.NodeKind(1): {1, 2, 3},
				},
			},
			kind:     ast.NodeKind(1),
			expected: 3,
		},
		{
			name: "PopValue from empty slice",
			ctx: &Context{
				values: map[ast.NodeKind][]any{
					ast.NodeKind(1): {},
				},
			},
			kind:     ast.NodeKind(1),
			expected: nil,
		},
		{
			name: "PopValue from nil map",
			ctx: &Context{
				values: nil,
			},
			kind:     ast.NodeKind(1),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.ctx.PopValue(tc.kind)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
