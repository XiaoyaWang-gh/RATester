package render

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPeekValue_1(t *testing.T) {
	type testCase struct {
		name     string
		ctx      *Context
		nodeKind ast.NodeKind
		want     any
	}

	testCases := []testCase{
		{
			name: "Returns the last value pushed for the given node kind",
			ctx: &Context{
				values: map[ast.NodeKind][]any{
					ast.NodeKind(1): {1, 2, 3},
				},
			},
			nodeKind: ast.NodeKind(1),
			want:     3,
		},
		{
			name: "Returns nil if the node kind has no values",
			ctx: &Context{
				values: map[ast.NodeKind][]any{
					ast.NodeKind(2): {4, 5, 6},
				},
			},
			nodeKind: ast.NodeKind(1),
			want:     nil,
		},
		{
			name:     "Returns nil if the context has no values",
			ctx:      &Context{},
			nodeKind: ast.NodeKind(1),
			want:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.ctx.PeekValue(tc.nodeKind)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PeekValue() = %v, want %v", got, tc.want)
			}
		})
	}
}
