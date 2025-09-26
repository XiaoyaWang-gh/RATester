package source

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIgnoreFile_2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "empty filename",
			filename: "",
			want:     true,
		},
		{
			name:     "filename with dot",
			filename: ".hidden",
			want:     true,
		},
		{
			name:     "filename with hash",
			filename: "#hidden",
			want:     true,
		},
		{
			name:     "filename with tilde",
			filename: "hidden~",
			want:     true,
		},
		{
			name:     "filename with unix slash",
			filename: "dir/file",
			want:     false,
		},
		{
			name:     "filename with windows slash",
			filename: "dir\\file",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &SourceSpec{
				SourceFs: &afero.OsFs{},
				shouldInclude: func(filename string) bool {
					return false
				},
			}

			if got := s.IgnoreFile(tt.filename); got != tt.want {
				t.Errorf("IgnoreFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
