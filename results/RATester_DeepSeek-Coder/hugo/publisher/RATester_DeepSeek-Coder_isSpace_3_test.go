package publisher

import (
	"fmt"
	"testing"
)

func TestIsSpace_3(t *testing.T) {
	testCases := []struct {
		name  string
		input byte
		want  bool
	}{
		{"space", ' ', true},
		{"tab", '\t', true},
		{"newline", '\n', true},
		{"non-space", 'a', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isSpace(tc.input)
			if got != tc.want {
				t.Errorf("isSpace(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
