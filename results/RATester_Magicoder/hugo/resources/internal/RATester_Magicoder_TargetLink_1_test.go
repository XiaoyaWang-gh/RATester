package internal

import (
	"fmt"
	"testing"
)

func TestTargetLink_1(t *testing.T) {
	type test struct {
		name     string
		input    ResourcePaths
		expected string
	}

	tests := []test{
		{
			name: "Test case 1",
			input: ResourcePaths{
				Dir:             "/dir1/dir2",
				BaseDirTarget:   "/base/target",
				BaseDirLink:     "/base/link",
				TargetBasePaths: []string{"/base/path1", "/base/path2"},
				File:            "data.json",
			},
			expected: "/base/target/dir1/dir2/data.json",
		},
		{
			name: "Test case 2",
			input: ResourcePaths{
				Dir:             "/dir1",
				BaseDirTarget:   "/base/target",
				BaseDirLink:     "/base/link",
				TargetBasePaths: []string{"/base/path1", "/base/path2"},
				File:            "data.json",
			},
			expected: "/base/target/dir1/data.json",
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.TargetLink()
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
