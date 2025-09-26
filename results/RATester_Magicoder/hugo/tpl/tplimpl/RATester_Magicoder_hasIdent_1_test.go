package tplimpl

import (
	"fmt"
	"testing"
)

func TesthasIdent_1(t *testing.T) {
	testCases := []struct {
		name   string
		idents []string
		ident  string
		want   bool
	}{
		{
			name:   "IdentExists",
			idents: []string{"foo", "bar", "baz"},
			ident:  "bar",
			want:   true,
		},
		{
			name:   "IdentDoesNotExist",
			idents: []string{"foo", "bar", "baz"},
			ident:  "qux",
			want:   false,
		},
		{
			name:   "EmptyIdentList",
			idents: []string{},
			ident:  "qux",
			want:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &templateContext{}
			got := ctx.hasIdent(tc.idents, tc.ident)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
