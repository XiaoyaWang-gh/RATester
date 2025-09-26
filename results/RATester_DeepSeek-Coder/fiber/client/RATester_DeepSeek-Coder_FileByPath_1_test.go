package client

import (
	"fmt"
	"testing"
)

func TestFileByPath_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		files    []*File
		path     string
		expected *File
	}

	testCases := []testCase{
		{
			name: "TestFileByPath_ExistingFile",
			files: []*File{
				{
					path: "/path/to/file1",
				},
				{
					path: "/path/to/file2",
				},
			},
			path: "/path/to/file1",
			expected: &File{
				path: "/path/to/file1",
			},
		},
		{
			name: "TestFileByPath_NonExistingFile",
			files: []*File{
				{
					path: "/path/to/file1",
				},
				{
					path: "/path/to/file2",
				},
			},
			path:     "/path/to/file3",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Request{
				files: tc.files,
			}

			result := r.FileByPath(tc.path)

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
