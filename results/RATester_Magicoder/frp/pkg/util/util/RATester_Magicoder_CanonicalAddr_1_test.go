package util

import (
	"fmt"
	"testing"
)

func TestCanonicalAddr_1(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		port     int
		expected string
	}{
		{
			name:     "standard http port",
			host:     "example.com",
			port:     80,
			expected: "example.com",
		},
		{
			name:     "standard https port",
			host:     "example.com",
			port:     443,
			expected: "example.com",
		},
		{
			name:     "non-standard port",
			host:     "example.com",
			port:     8080,
			expected: "example.com:8080",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			addr := CanonicalAddr(test.host, test.port)
			if addr != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, addr)
			}
		})
	}
}
