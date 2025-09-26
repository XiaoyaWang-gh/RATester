package template

import (
	"fmt"
	"testing"
)

func TestIsHTMLSpace_1(t *testing.T) {
	testCases := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test case 1", ' ', true},
		{"Test case 2", '\t', true},
		{"Test case 3", '\n', true},
		{"Test case 4", '\r', true},
		{"Test case 5", 'a', false},
		{"Test case 6", '0', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isHTMLSpace(tc.input)
			if got != tc.want {
				t.Errorf("Expected %t, but got %t", tc.want, got)
			}
		})
	}
}
