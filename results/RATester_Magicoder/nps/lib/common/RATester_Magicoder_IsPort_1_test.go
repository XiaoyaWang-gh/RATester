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
		{"valid port", "8080", true},
		{"invalid port", "0", false},
		{"invalid port", "65536", false},
		{"invalid port", "abc", false},
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
