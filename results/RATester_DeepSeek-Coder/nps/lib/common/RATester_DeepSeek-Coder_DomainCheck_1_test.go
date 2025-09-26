package common

import (
	"fmt"
	"testing"
)

func TestDomainCheck_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Test case 1", "http://www.google.com", true},
		{"Test case 2", "https://www.google.com", true},
		{"Test case 3", "www.google.com", true},
		{"Test case 4", "google.com", true},
		{"Test case 5", "http://google", false},
		{"Test case 6", "https://google", false},
		{"Test case 7", "google", false},
		{"Test case 8", "http://www.google.com/path", true},
		{"Test case 9", "https://www.google.com/path", true},
		{"Test case 10", "www.google.com/path", true},
		{"Test case 11", "google.com/path", true},
		{"Test case 12", "http://google/path", false},
		{"Test case 13", "https://google/path", false},
		{"Test case 14", "google/path", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DomainCheck(tt.args); got != tt.want {
				t.Errorf("DomainCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
