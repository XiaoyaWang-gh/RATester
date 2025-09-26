package http

import (
	"fmt"
	"testing"
)

func TestHasPort_1(t *testing.T) {
	tests := []struct {
		name string
		host string
		want bool
	}{
		{"no port", "example.com", false},
		{"with port", "example.com:80", true},
		{"IPv6 with port", "[2001:db8::1]:80", true},
		{"IPv6 no port", "[2001:db8::1]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := hasPort(tt.host); got != tt.want {
				t.Errorf("hasPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
