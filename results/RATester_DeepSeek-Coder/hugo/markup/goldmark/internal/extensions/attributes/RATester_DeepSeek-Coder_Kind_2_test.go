package attributes

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestKind_2(t *testing.T) {
	tests := []struct {
		name string
		want ast.NodeKind
	}{
		{
			name: "TestKind",
			want: kindAttributesBlock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			a := &attributesBlock{}
			if got := a.Kind(); got != tt.want {
				t.Errorf("attributesBlock.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}
