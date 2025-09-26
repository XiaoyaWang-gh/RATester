package doctree

import (
	"fmt"
	"testing"
)

func TestShape_1(t *testing.T) {
	tree := TreeShiftTree[int]{
		d:     1,
		v:     0,
		zero:  0,
		trees: []*SimpleTree[int]{},
	}

	tests := []struct {
		name string
		d    int
		v    int
		want *TreeShiftTree[int]
	}{
		{
			name: "Test case 1",
			d:    1,
			v:    0,
			want: &tree,
		},
		{
			name: "Test case 2",
			d:    2,
			v:    0,
			want: &tree,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tree.Shape(tt.d, tt.v)
			if got != tt.want {
				t.Errorf("Shape() = %v, want %v", got, tt.want)
			}
		})
	}
}
