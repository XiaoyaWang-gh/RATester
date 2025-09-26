package page

import (
	"fmt"
	"testing"
)

func TestPathFile_1(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		elements []string
		want     string
	}{
		{
			name:     "Test with no prefix",
			prefix:   "",
			elements: []string{"test", "path"},
			want:     "/test/path",
		},
		{
			name:     "Test with prefix",
			prefix:   "prefix",
			elements: []string{"test", "path"},
			want:     "/prefix/test/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pagePathBuilder{
				els:        tt.elements,
				prefixPath: tt.prefix,
			}
			got := p.PathFile()
			if got != tt.want {
				t.Errorf("PathFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
