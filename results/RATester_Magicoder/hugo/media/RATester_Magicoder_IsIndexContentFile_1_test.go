package media

import (
	"fmt"
	"testing"
)

func TestIsIndexContentFile_1(t *testing.T) {
	contentTypes := ContentTypes{
		HTML: Type{
			Type:        "text/html",
			MainType:    "text",
			SubType:     "html",
			Delimiter:   ".",
			SuffixesCSV: "html,htm",
		},
	}

	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "index.html",
			filename: "index.html",
			want:     true,
		},
		{
			name:     "_index.html",
			filename: "_index.html",
			want:     true,
		},
		{
			name:     "index.htm",
			filename: "index.htm",
			want:     true,
		},
		{
			name:     "_index.htm",
			filename: "_index.htm",
			want:     true,
		},
		{
			name:     "index.txt",
			filename: "index.txt",
			want:     false,
		},
		{
			name:     "_index.txt",
			filename: "_index.txt",
			want:     false,
		},
		{
			name:     "index.md",
			filename: "index.md",
			want:     false,
		},
		{
			name:     "_index.md",
			filename: "_index.md",
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

			if got := contentTypes.IsIndexContentFile(tt.filename); got != tt.want {
				t.Errorf("IsIndexContentFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
