package pageparser

import (
	"fmt"
	"testing"
)

func TestHasShortcode_1(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Test case 1",
			s:    "{{shortcode}}",
			want: true,
		},
		{
			name: "Test case 2",
			s:    "No shortcode",
			want: false,
		},
		{
			name: "Test case 3",
			s:    "{{shortcode}} and more text",
			want: true,
		},
		{
			name: "Test case 4",
			s:    "{{shortcode",
			want: false,
		},
		{
			name: "Test case 5",
			s:    "shortcode}}",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasShortcode(tt.s); got != tt.want {
				t.Errorf("HasShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
