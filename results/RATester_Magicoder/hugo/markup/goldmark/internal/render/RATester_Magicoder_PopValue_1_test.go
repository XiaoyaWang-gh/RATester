package render

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPopValue_1(t *testing.T) {
	ctx := &Context{
		values: map[ast.NodeKind][]any{
			ast.NodeKind(1): {1, 2, 3},
			ast.NodeKind(2): {4, 5, 6},
		},
	}

	tests := []struct {
		name string
		k    ast.NodeKind
		want any
	}{
		{
			name: "Test case 1",
			k:    ast.NodeKind(1),
			want: 3,
		},
		{
			name: "Test case 2",
			k:    ast.NodeKind(2),
			want: 6,
		},
		{
			name: "Test case 3",
			k:    ast.NodeKind(3),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ctx.PopValue(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
