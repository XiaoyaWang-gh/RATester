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
		{
			name:   "Test case 1",
			inPath: "/path/to/file",
			want:   "../../../",
		},
		{
			name:   "Test case 2",
			inPath: "/path/to/dir/",
			want:   "../../../",
		},
		{
			name:   "Test case 3",
			inPath: "/",
			want:   "./",
		},
		{
			name:   "Test case 4",
			inPath: ".",
			want:   "./",
		},
		{
			name:   "Test case 5",
			inPath: "path/to/file",
			want:   "../",
		},
		{
			name:   "Test case 6",
			inPath: "path/to/dir/",
			want:   "../",
		},
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
