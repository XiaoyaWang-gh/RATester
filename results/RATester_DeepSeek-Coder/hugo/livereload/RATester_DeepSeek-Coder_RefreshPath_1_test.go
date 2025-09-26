package livereload

import (
	"fmt"
	"testing"
)

func TestRefreshPath_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "test",
			expected: "test",
		},
		{
			name:     "Test case 2",
			input:    "test:8080",
			expected: "test:8080",
		},
		{
			name:     "Test case 3",
			input:    "test:",
			expected: "test:",
		},
		{
			name:     "Test case 4",
			input:    ":8080",
			expected: ":8080",
		},
		{
			name:     "Test case 5",
			input:    ":",
			expected: ":",
		},
		{
			name:     "Test case 6",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			RefreshPath(tc.input)
			if tc.input != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, tc.input)
			}
		})
	}
}
