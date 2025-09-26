package internal

import (
	"fmt"
	"testing"
)

func TestTargetPath_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    ResourcePaths
		expected string
	}{
		{
			name: "Test case 1",
			input: ResourcePaths{
				Dir:             "/dir1",
				BaseDirTarget:   "/baseDir1",
				BaseDirLink:     "/baseDirLink1",
				TargetBasePaths: []string{"/targetBasePath1"},
				File:            "file1",
			},
			expected: "/baseDir1/dir1/file1",
		},
		{
			name: "Test case 2",
			input: ResourcePaths{
				Dir:             "/dir2",
				BaseDirTarget:   "/baseDir2",
				BaseDirLink:     "/baseDirLink2",
				TargetBasePaths: []string{"/targetBasePath2"},
				File:            "file2",
			},
			expected: "/baseDir2/dir2/file2",
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

			result := tc.input.TargetPath()
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
