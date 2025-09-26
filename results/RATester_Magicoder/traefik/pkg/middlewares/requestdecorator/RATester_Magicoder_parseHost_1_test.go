package requestdecorator

import (
	"fmt"
	"testing"
)

func TestParseHost_1(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{"no port", "localhost", "localhost"},
		{"with port", "localhost:8080", "localhost"},
		{"ipv6 with port", "[::1]:8080", "::1"},
		{"ipv6 no port", "[::1]", "::1"},
		{"invalid ipv6", "[::1", "[::1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseHost(tt.addr); got != tt.want {
				t.Errorf("parseHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
