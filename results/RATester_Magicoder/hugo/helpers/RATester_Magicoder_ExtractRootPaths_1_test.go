package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractRootPaths_1(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  []string
	}{
		{
			name:  "empty paths",
			paths: []string{},
			want:  []string{},
		},
		{
			name:  "single path",
			paths: []string{"/path/to/file"},
			want:  []string{"path"},
		},
		{
			name:  "multiple paths",
			paths: []string{"/path/to/file", "/another/path/to/file"},
			want:  []string{"path", "another"},
		},
		{
			name:  "paths with empty sections",
			paths: []string{"/path//to/file", "/another//path/to/file"},
			want:  []string{"path", "another"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ExtractRootPaths(tt.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractRootPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
