package goldmark

import (
	"fmt"
	"testing"
)

func TestSanitizeAnchorName_2(t *testing.T) {
	converter := &goldmarkConverter{
		sanitizeAnchorName: func(s string) string {
			return "test"
		},
	}

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple string",
			input:    "test",
			expected: "test",
		},
		{
			name:     "string with special characters",
			input:    "test@#$%^&*()_+`~-=[]\\\";',./{}|\":<>?test",
			expected: "test",
		},
		{
			name:     "string with spaces",
			input:    "test with spaces",
			expected: "test with spaces",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := converter.SanitizeAnchorName(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
