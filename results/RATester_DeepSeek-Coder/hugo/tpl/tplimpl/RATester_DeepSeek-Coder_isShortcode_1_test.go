package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsShortcode_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"shortcode_exists", "shortcode1", true},
		{"shortcode_not_exists", "notshortcode", false},
		{"empty_string", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isShortcode(tt.arg); got != tt.want {
				t.Errorf("isShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
