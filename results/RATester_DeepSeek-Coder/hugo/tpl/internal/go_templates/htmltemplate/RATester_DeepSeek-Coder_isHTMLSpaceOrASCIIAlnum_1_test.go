package template

import (
	"fmt"
	"testing"
)

func TestIsHTMLSpaceOrASCIIAlnum_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    byte
		expected bool
	}{
		{"Test case 1", ' ', true},
		{"Test case 2", '\t', true},
		{"Test case 3", '\n', true},
		{"Test case 4", 'a', true},
		{"Test case 5", 'z', true},
		{"Test case 6", 'A', true},
		{"Test case 7", 'Z', true},
		{"Test case 8", '0', true},
		{"Test case 9", '9', true},
		{"Test case 10", '~', false},
		{"Test case 11", 0x80, false},
		{"Test case 12", 0x7F, true},
		{"Test case 13", 0x08, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isHTMLSpaceOrASCIIAlnum(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
