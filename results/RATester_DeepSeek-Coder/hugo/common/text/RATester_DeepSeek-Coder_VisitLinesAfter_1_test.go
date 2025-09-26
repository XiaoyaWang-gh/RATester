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
			name:     "single line",
			input:    "Hello, World!\n",
			expected: []string{"Hello, World!\n"},
		},
		{
			name:     "multiple lines",
			input:    "Hello, World!\nHow are you?\nI'm fine, thank you.\n",
			expected: []string{"Hello, World!\n", "How are you?\n", "I'm fine, thank you.\n"},
		},
		{
			name:     "no newline",
			input:    "Hello, World!\nHow are you?I'm fine, thank you.",
			expected: []string{"Hello, World!\nHow are you?I'm fine, thank you."},
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
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
