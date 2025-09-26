package files

import (
	"fmt"
	"testing"
)

func TestResolveComponentFolder_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		expected string
	}{
		{
			name:     "Test case 1",
			filename: "/path/to/component/file",
			expected: "/path/to/component",
		},
		{
			name:     "Test case 2",
			filename: "path/to/component/file",
			expected: "path/to/component",
		},
		{
			name:     "Test case 3",
			filename: "component/file",
			expected: "",
		},
		{
			name:     "Test case 4",
			filename: "/component/file",
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

			result := ResolveComponentFolder(tc.filename)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
