package install

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMkidrDirAll_1(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		v        []string
		expected error
	}{
		{
			name:     "Test case 1",
			path:     "/tmp",
			v:        []string{"test1", "test2"},
			expected: nil,
		},
		{
			name:     "Test case 2",
			path:     "/tmp/test1",
			v:        []string{"test3", "test4"},
			expected: nil,
		},
		{
			name:     "Test case 3",
			path:     "/tmp/test1/test3",
			v:        []string{"test5", "test6"},
			expected: nil,
		},
		{
			name:     "Test case 4",
			path:     "/tmp/test1/test3/test5",
			v:        []string{"test7", "test8"},
			expected: nil,
		},
		{
			name:     "Test case 5",
			path:     "/tmp/test1/test3/test5/test7",
			v:        []string{"test9", "test10"},
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

			MkidrDirAll(tc.path, tc.v...)
			for _, item := range tc.v {
				_, err := os.Stat(filepath.Join(tc.path, item))
				if err != nil {
					t.Errorf("Expected error to be nil, got %v", err)
				}
			}
		})
	}
}
