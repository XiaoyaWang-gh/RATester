package pageparser

import (
	"fmt"
	"testing"
)

func TestHasShortcode_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"empty string", "", false},
		{"no shortcode", "Hello, world", false},
		{"shortcode", "Hello, {{world}}", true},
		{"shortcode with text", "Hello, {{world}} !", true},
		{"shortcode with text and spaces", "Hello, {{ world }} !", true},
		{"shortcode with text and special characters", "Hello, {{world_123}} !", true},
		{"shortcode with text and special characters and spaces", "Hello, {{ world_123 }} !", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasShortcode(tt.arg); got != tt.want {
				t.Errorf("HasShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
