package template

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestJsIsSpecial_1(t *testing.T) {
	testCases := []struct {
		name  string
		input rune
		want  bool
	}{
		{"Test case 1", '\\', true},
		{"Test case 2", '\'', true},
		{"Test case 3", '"', true},
		{"Test case 4", '<', true},
		{"Test case 5", '>', true},
		{"Test case 6", '&', true},
		{"Test case 7", '=', true},
		{"Test case 8", ' ', false},
		{"Test case 9", 'a', false},
		{"Test case 10", 'z', false},
		{"Test case 11", '0', false},
		{"Test case 12", '9', false},
		{"Test case 13", utf8.MaxRune + 1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := jsIsSpecial(tc.input)
			if got != tc.want {
				t.Errorf("jsIsSpecial(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
