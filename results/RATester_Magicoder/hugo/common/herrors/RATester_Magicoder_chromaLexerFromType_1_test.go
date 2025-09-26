package herrors

import (
	"fmt"
	"testing"
)

func TestchromaLexerFromType_1(t *testing.T) {
	tests := []struct {
		name     string
		fileType string
		want     string
	}{
		{
			name:     "Test HTML",
			fileType: "html",
			want:     "go-html-template",
		},
		{
			name:     "Test HTM",
			fileType: "htm",
			want:     "go-html-template",
		},
		{
			name:     "Test Other",
			fileType: "other",
			want:     "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := chromaLexerFromType(tt.fileType); got != tt.want {
				t.Errorf("chromaLexerFromType() = %v, want %v", got, tt.want)
			}
		})
	}
}
