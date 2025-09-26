package utils

import (
	"fmt"
	"testing"
)

func TestneedsEncoding_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"Test case 1", "Hello, World!", false},
		{"Test case 2", "Hello, \tWorld!", true},
		{"Test case 3", "Hello, \nWorld!", true},
		{"Test case 4", "Hello, \rWorld!", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := needsEncoding(tt.s); got != tt.want {
				t.Errorf("needsEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
