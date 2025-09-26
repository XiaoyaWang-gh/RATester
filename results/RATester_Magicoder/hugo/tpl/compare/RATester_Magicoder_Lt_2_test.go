package compare

import (
	"fmt"
	"testing"
)

func TestLt_2(t *testing.T) {
	n := &Namespace{}

	tests := []struct {
		name   string
		first  any
		others []any
		want   bool
	}{
		{
			name:   "Test case 1",
			first:  1,
			others: []any{2, 3, 4},
			want:   true,
		},
		{
			name:   "Test case 2",
			first:  5,
			others: []any{2, 3, 4},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  "a",
			others: []any{"b", "c", "d"},
			want:   true,
		},
		{
			name:   "Test case 4",
			first:  "e",
			others: []any{"b", "c", "d"},
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

			if got := n.Lt(tt.first, tt.others...); got != tt.want {
				t.Errorf("Lt() = %v, want %v", got, tt.want)
			}
		})
	}
}
