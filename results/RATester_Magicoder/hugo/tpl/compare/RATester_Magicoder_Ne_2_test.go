package compare

import (
	"fmt"
	"testing"
)

func TestNe_2(t *testing.T) {
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
			first:  1,
			others: []any{1, 2, 3},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  "hello",
			others: []any{"world", "golang"},
			want:   true,
		},
		{
			name:   "Test case 4",
			first:  "hello",
			others: []any{"hello", "world", "golang"},
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

			if got := n.Ne(tt.first, tt.others...); got != tt.want {
				t.Errorf("Ne() = %v, want %v", got, tt.want)
			}
		})
	}
}
