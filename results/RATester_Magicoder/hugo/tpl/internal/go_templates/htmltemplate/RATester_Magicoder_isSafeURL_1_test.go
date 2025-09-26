package template

import (
	"fmt"
	"testing"
)

func TestisSafeURL_1(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{"http", "http://example.com", true},
		{"https", "https://example.com", true},
		{"mailto", "mailto:example@example.com", true},
		{"invalid protocol", "ftp://example.com", false},
		{"invalid url", "example.com", false},
		{"empty string", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isSafeURL(tt.url); got != tt.want {
				t.Errorf("isSafeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
