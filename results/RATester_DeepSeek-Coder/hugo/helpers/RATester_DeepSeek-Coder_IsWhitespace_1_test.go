package helpers

import (
	"fmt"
	"testing"
)

func TestIsWhitespace_1(t *testing.T) {
	testCases := []struct {
		name  string
		input rune
		want  bool
	}{
		{"space", ' ', true},
		{"tab", '\t', true},
		{"newline", '\n', true},
		{"return", '\r', true},
		{"non-whitespace", 'a', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := IsWhitespace(tc.input)
			if got != tc.want {
				t.Errorf("IsWhitespace(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
