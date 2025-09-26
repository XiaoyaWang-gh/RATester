package common

import (
	"fmt"
	"testing"
)

func TestDomainCheck_1(t *testing.T) {
	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{"Test case 1", "google.com", true},
		{"Test case 2", "http://google.com", true},
		{"Test case 3", "https://google.com", true},
		{"Test case 4", "http://google.com/", true},
		{"Test case 5", "https://google.com/", true},
		{"Test case 6", "google", false},
		{"Test case 7", "http://google", false},
		{"Test case 8", "https://google", false},
		{"Test case 9", "http://google/", false},
		{"Test case 10", "https://google/", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DomainCheck(tt.domain); got != tt.want {
				t.Errorf("DomainCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
