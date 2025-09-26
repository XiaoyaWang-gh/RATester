package template

import (
	"fmt"
	"testing"
)

func TestIsHex_1(t *testing.T) {
	testCases := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test Case 1", 'a', true},
		{"Test Case 2", 'f', true},
		{"Test Case 3", 'F', true},
		{"Test Case 4", '0', true},
		{"Test Case 5", '9', true},
		{"Test Case 6", 'g', false},
		{"Test Case 7", 'G', false},
		{"Test Case 8", 'x', false},
		{"Test Case 9", 'X', false},
		{"Test Case 10", ' ', false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isHex(tc.input)
			if got != tc.want {
				t.Errorf("Expected %t, but got %t", tc.want, got)
			}
		})
	}
}
