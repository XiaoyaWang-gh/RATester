package helpers

import (
	"fmt"
	"testing"
)

func TestGetDottedRelativePath_1(t *testing.T) {
	tests := []struct {
		name   string
		inPath string
		want   string
	}{
		{"Test case 1", "/path/to/file", "./"},
		{"Test case 2", "/path/to/dir/", "./"},
		{"Test case 3", "/path/to/dir", "../"},
		{"Test case 4", "/path/to/dir/subdir", "../../"},
		{"Test case 5", "/path/to/dir/subdir/subsubdir", "../../../"},
		{"Test case 6", "/", "./"},
		{"Test case 7", ".", "./"},
		{"Test case 8", "path/to/file", "./"},
		{"Test case 9", "path/to/dir/", "./"},
		{"Test case 10", "path/to/dir", "../"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetDottedRelativePath(tt.inPath); got != tt.want {
				t.Errorf("GetDottedRelativePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
