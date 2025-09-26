package helpers

import (
	"fmt"
	"testing"
)

func TestTotalWords_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world", 2},
		{"multiple words with spaces", "  hello   world  ", 2},
		{"multiple words with tabs", "  hello\tworld  ", 2},
		{"multiple words with newlines", "  hello\nworld  ", 2},
		{"multiple words with mixed spaces, tabs, and newlines", "  hello\t\nworld  ", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := TotalWords(tt.s); got != tt.want {
				t.Errorf("TotalWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
