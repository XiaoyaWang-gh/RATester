package herrors

import (
	"fmt"
	"testing"
)

func TestchromaLexerFromFilename_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "Test with filename containing 'layouts'",
			filename: "layouts/test.html",
			want:     "go-html-template",
		},
		{
			name:     "Test with filename not containing 'layouts'",
			filename: "test.go",
			want:     "go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := chromaLexerFromFilename(tt.filename); got != tt.want {
				t.Errorf("chromaLexerFromFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
