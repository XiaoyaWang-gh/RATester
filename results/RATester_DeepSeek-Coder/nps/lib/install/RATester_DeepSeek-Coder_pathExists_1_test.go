package install

import (
	"errors"
	"fmt"
	"testing"
)

func TestPathExists_1(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
		err      error
	}{
		{
			name:     "Test Case 1: Path Exists",
			path:     "/tmp",
			expected: true,
			err:      nil,
		},
		{
			name:     "Test Case 2: Path Does Not Exist",
			path:     "/tmp/nonexistent",
			expected: false,
			err:      nil,
		},
		{
			name:     "Test Case 3: Error Case",
			path:     "/",
			expected: false,
			err:      errors.New("permission denied"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			exists, err := pathExists(tc.path)
			if exists != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, exists)
			}
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
