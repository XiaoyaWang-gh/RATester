package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsBaseTemplatePath_1(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "Test case 1",
			path:     "/path/to/baseFileBase",
			expected: true,
		},
		{
			name:     "Test case 2",
			path:     "/path/to/anotherFile",
			expected: false,
		},
		{
			name:     "Test case 3",
			path:     "/path/to/baseFileBase.txt",
			expected: true,
		},
		{
			name:     "Test case 4",
			path:     "/path/to/baseFileBase/subdir",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isBaseTemplatePath(tc.path)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
