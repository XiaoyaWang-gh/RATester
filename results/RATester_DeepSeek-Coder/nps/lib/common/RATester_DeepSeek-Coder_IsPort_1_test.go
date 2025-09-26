package common

import (
	"fmt"
	"testing"
)

func TestIsPort_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test case 1", "8080", true},
		{"Test case 2", "65536", true},
		{"Test case 3", "0", false},
		{"Test case 4", "65537", false},
		{"Test case 5", "abc", false},
		{"Test case 6", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsPort(tt.arg); got != tt.want {
				t.Errorf("IsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
