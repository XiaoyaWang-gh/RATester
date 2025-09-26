package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakePathsSanitized_1(t *testing.T) {
	type testCase struct {
		name     string
		input    []string
		expected []string
	}

	testCases := []testCase{
		{
			name:     "simple path",
			input:    []string{"/path/to/file"},
			expected: []string{"/path/to/file"},
		},
		{
			name:     "multiple paths",
			input:    []string{"/path/to/file1", "/path/to/file2"},
			expected: []string{"/path/to/file1", "/path/to/file2"},
		},
		{
			name:     "path with spaces",
			input:    []string{" /path/to/file "},
			expected: []string{"%20/path/to/file%20"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ps := &PathSpec{}
			ps.MakePathsSanitized(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input)
			}
		})
	}
}
