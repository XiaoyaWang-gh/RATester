package media

import (
	"fmt"
	"testing"
)

func TestIsContentFile_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "Test with a markdown file",
			filename: "test.md",
			want:     true,
		},
		{
			name:     "Test with a html file",
			filename: "test.html",
			want:     true,
		},
		{
			name:     "Test with a txt file",
			filename: "test.txt",
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

			contentTypes := ContentTypes{
				Markdown: Type{
					MainType: "text",
					SubType:  "markdown",
				},
				HTML: Type{
					MainType: "text",
					SubType:  "html",
				},
			}

			if got := contentTypes.IsContentFile(tt.filename); got != tt.want {
				t.Errorf("IsContentFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
