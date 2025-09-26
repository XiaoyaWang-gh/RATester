package compare

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLtCollate_1(t *testing.T) {
	type testStruct struct {
		name     string
		collator *langs.Collator
		first    any
		others   []any
		want     bool
	}

	tests := []testStruct{
		{
			name:     "TestLtCollate_True",
			collator: &langs.Collator{},
			first:    "a",
			others:   []any{"b", "c", "d"},
			want:     true,
		},
		{
			name:     "TestLtCollate_False",
			collator: &langs.Collator{},
			first:    "a",
			others:   []any{"a", "b", "c"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			n := &Namespace{}
			got := n.LtCollate(tt.collator, tt.first, tt.others...)
			if got != tt.want {
				t.Errorf("LtCollate() = %v, want %v", got, tt.want)
			}
		})
	}
}
