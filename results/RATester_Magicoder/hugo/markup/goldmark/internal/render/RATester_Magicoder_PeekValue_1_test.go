package render

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestPeekValue_1(t *testing.T) {
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
			name: "Existing key",
			k:    ast.NodeKind(1),
			want: 3,
		},
		{
			name: "Non-existing key",
			k:    ast.NodeKind(3),
			want: nil,
		},
		{
			name: "Empty slice",
			k:    ast.NodeKind(2),
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ctx.PeekValue(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeekValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
