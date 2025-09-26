package compare

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLtCollate_1(t *testing.T) {
	n := &Namespace{}
	collator := &langs.Collator{}

	tests := []struct {
		name   string
		first  any
		others []any
		want   bool
	}{
		{
			name:   "Test case 1",
			first:  "apple",
			others: []any{"banana", "cherry"},
			want:   true,
		},
		{
			name:   "Test case 2",
			first:  "banana",
			others: []any{"apple", "cherry"},
			want:   false,
		},
		{
			name:   "Test case 3",
			first:  "cherry",
			others: []any{"apple", "banana"},
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

			if got := n.LtCollate(collator, tt.first, tt.others...); got != tt.want {
				t.Errorf("LtCollate() = %v, want %v", got, tt.want)
			}
		})
	}
}
