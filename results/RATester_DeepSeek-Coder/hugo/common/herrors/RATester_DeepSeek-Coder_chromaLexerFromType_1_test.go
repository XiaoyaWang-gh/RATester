package herrors

import (
	"fmt"
	"testing"
)

func TestChromaLexerFromType_1(t *testing.T) {
	tests := []struct {
		name     string
		fileType string
		want     string
	}{
		{"HTML", "html", "go-html-template"},
		{"HTM", "htm", "go-html-template"},
		{"Unknown", "unknown", "unknown"},
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
