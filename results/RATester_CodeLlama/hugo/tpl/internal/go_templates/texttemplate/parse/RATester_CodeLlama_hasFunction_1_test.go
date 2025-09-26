package parse

import (
	"fmt"
	"testing"
)

func TestHasFunction_1(t *testing.T) {
	tests := []struct {
		name string
		t    *Tree
		want bool
	}{
		{
			name: "test hasFunction",
			t: &Tree{
				funcs: []map[string]any{
					{
						"name": nil,
					},
				},
			},
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

			if got := tt.t.hasFunction(tt.name); got != tt.want {
				t.Errorf("Tree.hasFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
