package text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVisitLinesAfter_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single line",
			input:    "Hello, world!",
			expected: []string{"Hello, world!"},
		},
		{
			name:     "multiple lines",
			input:    "Hello,\nworld!\nThis is a test.",
			expected: []string{"Hello,\n", "world!\n", "This is a test."},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var actual []string
			VisitLinesAfter(tc.input, func(line string) {
				actual = append(actual, line)
			})

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, actual)
			}
		})
	}
}
