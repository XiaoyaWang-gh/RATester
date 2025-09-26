package common

import (
	"fmt"
	"testing"
)

func TestFormatAddress_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"IPv4", "8080", "127.0.0.1:8080"},
		{"IPv6", "[2001:db8::68]", "[2001:db8::68]"},
		{"IPv6 with port", "[2001:db8::68]:8080", "[2001:db8::68]:8080"},
		{"localhost", "localhost", "127.0.0.1:localhost"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatAddress(tt.arg); got != tt.want {
				t.Errorf("FormatAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
