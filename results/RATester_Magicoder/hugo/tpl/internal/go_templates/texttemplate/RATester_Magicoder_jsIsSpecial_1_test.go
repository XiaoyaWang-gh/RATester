package template

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestjsIsSpecial_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Test case 1", '\\', true},
		{"Test case 2", '\'', true},
		{"Test case 3", '"', true},
		{"Test case 4", '<', true},
		{"Test case 5", '>', true},
		{"Test case 6", '&', true},
		{"Test case 7", '=', true},
		{"Test case 8", ' ', false},
		{"Test case 9", 'a', false},
		{"Test case 10", 'A', false},
		{"Test case 11", '1', false},
		{"Test case 12", '@', false},
		{"Test case 13", '~', false},
		{"Test case 14", '', false},
		{"Test case 15", utf8.RuneSelf, false},
		{"Test case 16", utf8.RuneSelf + 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := jsIsSpecial(tt.r); got != tt.want {
				t.Errorf("jsIsSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
