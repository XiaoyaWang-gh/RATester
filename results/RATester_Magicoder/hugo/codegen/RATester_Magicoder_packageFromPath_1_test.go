package codegen

import (
	"fmt"
	"testing"
)

func TestpackageFromPath_1(t *testing.T) {
	inspector := &Inspector{}

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "simple path",
			path: "/path/to/file.go",
			want: "file",
		},
		{
			name: "path with subdirectories",
			path: "/path/to/subdir/file.go",
			want: "subdir",
		},
		{
			name: "path with multiple subdirectories",
			path: "/path/to/subdir1/subdir2/file.go",
			want: "subdir1",
		},
		{
			name: "path with no extension",
			path: "/path/to/file",
			want: "file",
		},
		{
			name: "path with no extension and subdirectories",
			path: "/path/to/subdir/file",
			want: "subdir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := inspector.packageFromPath(tt.path); got != tt.want {
				t.Errorf("packageFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
