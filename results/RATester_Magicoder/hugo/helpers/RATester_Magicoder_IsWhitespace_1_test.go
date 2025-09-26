package helpers

import (
	"fmt"
	"testing"
)

func TestIsWhitespace_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"space", ' ', true},
		{"tab", '\t', true},
		{"newline", '\n', true},
		{"carriage return", '\r', true},
		{"non-whitespace", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsWhitespace(tt.r); got != tt.want {
				t.Errorf("IsWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
