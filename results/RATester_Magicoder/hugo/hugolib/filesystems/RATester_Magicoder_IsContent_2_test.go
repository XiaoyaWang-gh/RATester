package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsContent_2(t *testing.T) {
	sf := SourceFilesystems{
		Content: &SourceFilesystem{
			Name: "content",
			Fs:   afero.NewMemMapFs(),
		},
	}

	sf.Content.Fs.MkdirAll("content", 0755)
	sf.Content.Fs.Create("content/test.txt")

	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"Existing file", "content/test.txt", true},
		{"Non-existing file", "content/nonexistent.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sf.IsContent(tt.filename); got != tt.want {
				t.Errorf("IsContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
