package tpl

import (
	"fmt"
	"testing"
)

func TestStripHTML_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No HTML",
			input:    "This is a test",
			expected: "This is a test",
		},
		{
			name:     "Simple HTML",
			input:    "<p>This is a test</p>",
			expected: "This is a test",
		},
		{
			name:     "Complex HTML",
			input:    "<div><p>This is a <b>test</b></p></div>",
			expected: "This is a test",
		},
		{
			name:     "HTML with newline",
			input:    "<div><p>This is a test<br>on a new line</p></div>",
			expected: "This is a test on a new line",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := StripHTML(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
