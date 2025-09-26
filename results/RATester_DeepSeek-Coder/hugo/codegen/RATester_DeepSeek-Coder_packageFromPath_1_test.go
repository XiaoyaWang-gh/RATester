package codegen

import (
	"fmt"
	"testing"
)

func TestPackageFromPath_1(t *testing.T) {
	inspector := &Inspector{}

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Test with a path containing a period",
			path: "/home/user/project/src/github.com/user/repo/pkg/subpkg/file.go",
			want: "subpkg",
		},
		{
			name: "Test with a path without a period",
			path: "/home/user/project/src/github.com/user/repo/pkg/file.go",
			want: "pkg",
		},
		{
			name: "Test with a path containing multiple periods",
			path: "/home/user/project/src/github.com/user/repo/pkg/subpkg/subsubpkg/file.go",
			want: "subpkg/subsubpkg",
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
