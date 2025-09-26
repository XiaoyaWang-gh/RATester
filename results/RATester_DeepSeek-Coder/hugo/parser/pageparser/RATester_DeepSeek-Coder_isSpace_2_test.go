package pageparser

import (
	"fmt"
	"testing"
)

func TestIsSpace_2(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"space", ' ', true},
		{"tab", '\t', true},
		{"newline", '\n', false},
		{"a", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isSpace(tt.r); got != tt.want {
				t.Errorf("isSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
