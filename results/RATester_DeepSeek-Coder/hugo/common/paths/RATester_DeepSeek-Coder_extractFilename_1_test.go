package paths

import (
	"fmt"
	"testing"
)

func TestExtractFilename_1(t *testing.T) {

	testCases := []struct {
		name     string
		in       string
		ext      string
		base     string
		pathSep  string
		expected string
	}{
		{"Test case 1", "/path/to/file.txt", ".txt", "file.txt", "/", "file"},
		{"Test case 2", "/path/to/file.txt", ".txt", "file.txt", "\\", "file"},
		{"Test case 3", "/path/to/file.txt", ".txt", "file.txt", "", "file"},
		{"Test case 4", "/path/to/file.txt", "", "file.txt", "/", "file"},
		{"Test case 5", "/path/to/file.txt", ".txt", "", "/", ""},
		{"Test case 6", "/path/to/file.txt", ".txt", ".", "", ""},
		{"Test case 7", "/path/to/file.txt", ".txt", "..", "/", ""},
		{"Test case 8", "/path/to/file.txt", ".txt", "/", "/", ""},
		{"Test case 9", "/path/to/file.txt", ".txt", "\\", "\\", ""},
		{"Test case 10", "/path/to/file.txt", ".txt", "file", "/", "file"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := extractFilename(tc.in, tc.ext, tc.base, tc.pathSep)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
