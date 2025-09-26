package compare

import (
	"fmt"
	"testing"
)

func TestGt_2(t *testing.T) {
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
			others: []any{5, 7, 9},
			want:   true,
		},
		{
			name:   "Test case 2",
			first:  10,
			others: []any{15, 17, 19},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  10,
			others: []any{10, 10, 10},
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

			if got := n.Gt(tt.first, tt.others...); got != tt.want {
				t.Errorf("Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}
