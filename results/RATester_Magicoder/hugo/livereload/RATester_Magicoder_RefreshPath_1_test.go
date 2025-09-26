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
			input:    "test1",
			expected: "test1",
		},
		{
			name:     "Test case 2",
			input:    "test2",
			expected: "test2",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			RefreshPath(tc.input)
			// Add assertions to check the expected behavior
		})
	}
}
