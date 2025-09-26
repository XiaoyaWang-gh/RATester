package paths

import (
	"fmt"
	"testing"
)

func TestToSlashPreserveLeading_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "/"},
		{"root directory", "/", "/"},
		{"single file", "file.txt", "/file.txt"},
		{"directory", "dir/", "/dir"},
		{"nested files", "dir/file.txt", "/dir/file.txt"},
		{"absolute path", "/abs/path/file.txt", "/abs/path/file.txt"},
		{"windows path", "C:\\Windows\\file.txt", "/C:/Windows/file.txt"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ToSlashPreserveLeading(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
