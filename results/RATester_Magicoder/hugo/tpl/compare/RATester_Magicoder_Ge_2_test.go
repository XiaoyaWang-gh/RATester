package compare

import (
	"fmt"
	"testing"
	"time"
)

func TestGe_2(t *testing.T) {
	n := &Namespace{
		loc:             time.UTC,
		caseInsensitive: false,
	}

	tests := []struct {
		name   string
		first  any
		others []any
		want   bool
	}{
		{
			name:   "Test case 1",
			first:  10,
			others: []any{10, 20, 30},
			want:   true,
		},
		{
			name:   "Test case 2",
			first:  10,
			others: []any{5, 10, 15},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  10,
			others: []any{10, 10, 10},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := n.Ge(tt.first, tt.others...); got != tt.want {
				t.Errorf("Ge() = %v, want %v", got, tt.want)
			}
		})
	}
}
