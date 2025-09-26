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
			name:  "Test case 1",
			paths: []string{"/usr/local/bin", "/bin", "/usr/bin"},
			want:  []string{"usr/local/bin", "bin", "usr/bin"},
		},
		{
			name:  "Test case 2",
			paths: []string{"/", "/usr", "/usr/local"},
			want:  []string{"", "usr", "usr/local"},
		},
		{
			name:  "Test case 3",
			paths: []string{"/a/b/c", "/d/e/f", "/g/h/i"},
			want:  []string{"a/b/c", "d/e/f", "g/h/i"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ExtractRootPaths(tt.paths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractRootPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
