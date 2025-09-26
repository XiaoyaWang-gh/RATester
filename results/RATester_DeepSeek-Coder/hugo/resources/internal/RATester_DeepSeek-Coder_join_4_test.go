package internal

import (
	"fmt"
	"testing"
)

func TestJoin_4(t *testing.T) {
	type test struct {
		name     string
		input    ResourcePaths
		args     []string
		expected string
	}

	tests := []test{
		{
			name: "Test join with empty strings",
			input: ResourcePaths{
				Dir:             "/dir",
				BaseDirTarget:   "/base",
				BaseDirLink:     "/link",
				TargetBasePaths: []string{"/base1", "/base2"},
				File:            "file",
			},
			args:     []string{"", "", "dir2", ""},
			expected: "/dir/dir2/file",
		},
		{
			name: "Test join with non-empty strings",
			input: ResourcePaths{
				Dir:             "/dir",
				BaseDirTarget:   "/base",
				BaseDirLink:     "/link",
				TargetBasePaths: []string{"/base1", "/base2"},
				File:            "file",
			},
			args:     []string{"dir1", "dir2", "dir3", "dir4"},
			expected: "/dir/dir1/dir2/dir3/dir4/file",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.join(tc.args...)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
