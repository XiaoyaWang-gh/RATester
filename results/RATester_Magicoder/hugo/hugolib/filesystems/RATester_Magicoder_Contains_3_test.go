package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestContains_3(t *testing.T) {
	testCases := []struct {
		name     string
		files    []string
		filename string
		expected bool
	}{
		{
			name:     "file exists",
			files:    []string{"/path/to/file1", "/path/to/file2"},
			filename: "/path/to/file1",
			expected: true,
		},
		{
			name:     "file does not exist",
			files:    []string{"/path/to/file1", "/path/to/file2"},
			filename: "/path/to/file3",
			expected: false,
		},
		{
			name:     "empty file list",
			files:    []string{},
			filename: "/path/to/file1",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fs := &SourceFilesystem{
				Name:     "test",
				SourceFs: afero.NewMemMapFs(),
			}

			for _, file := range tc.files {
				fs.SourceFs.Create(file)
			}

			result := fs.Contains(tc.filename)

			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
