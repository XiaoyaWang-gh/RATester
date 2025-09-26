package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsStatic_2(t *testing.T) {
	sfs := SourceFilesystems{
		Static: map[string]*SourceFilesystem{
			"": {
				Fs: afero.NewMemMapFs(),
			},
		},
	}

	afero.WriteFile(sfs.Static[""].Fs, "test.txt", []byte("test"), 0644)

	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "Existing file",
			filename: "test.txt",
			want:     true,
		},
		{
			name:     "Non-existing file",
			filename: "nonexistent.txt",
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

			if got := sfs.IsStatic(tt.filename); got != tt.want {
				t.Errorf("IsStatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
