package compare

import (
	"fmt"
	"testing"
)

func TestLe_2(t *testing.T) {
	n := &Namespace{}

	tests := []struct {
		name   string
		first  any
		others []any
		want   bool
	}{
		{
			name:   "Test case 1",
			first:  10,
			others: []any{20, 30, 40},
			want:   true,
		},
		{
			name:   "Test case 2",
			first:  50,
			others: []any{40, 30, 20},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  "apple",
			others: []any{"banana", "cherry", "date"},
			want:   true,
		},
		{
			name:   "Test case 4",
			first:  "grape",
			others: []any{"kiwi", "mango", "orange"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := n.Le(tt.first, tt.others...); got != tt.want {
				t.Errorf("Le() = %v, want %v", got, tt.want)
			}
		})
	}
}
