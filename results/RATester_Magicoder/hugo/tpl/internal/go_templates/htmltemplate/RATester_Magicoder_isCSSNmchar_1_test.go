package template

import (
	"fmt"
	"testing"
)

func TestisCSSNmchar_1(t *testing.T) {

	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"lowercase", 'a', true},
		{"uppercase", 'A', true},
		{"digit", '0', true},
		{"dash", '-', true},
		{"underscore", '_', true},
		{"non-ascii-low", 0x80, true},
		{"non-ascii-high", 0xd7ff, true},
		{"non-ascii-low-2", 0xe000, true},
		{"non-ascii-high-2", 0xfffd, true},
		{"non-ascii-low-3", 0x10000, true},
		{"non-ascii-high-3", 0x10ffff, true},
		{"control", 8, false},
		{"non-ascii-char", '×', false},
		{"non-ascii-char-2", 'à', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isCSSNmchar(tt.r); got != tt.want {
				t.Errorf("isCSSNmchar() = %v, want %v", got, tt.want)
			}
		})
	}
}
