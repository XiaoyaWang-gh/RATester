package accesslog

import (
	"fmt"
	"testing"
)

func TestSilentSplitHostPort_1(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected struct {
			host string
			port string
		}
	}{
		{
			name:  "valid host and port",
			value: "localhost:8080",
			expected: struct {
				host string
				port string
			}{
				host: "localhost",
				port: "8080",
			},
		},
		{
			name:  "invalid host and port",
			value: "localhost",
			expected: struct {
				host string
				port string
			}{
				host: "localhost",
				port: "-",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			host, port := silentSplitHostPort(test.value)
			if host != test.expected.host || port != test.expected.port {
				t.Errorf("Expected (%s, %s), got (%s, %s)", test.expected.host, test.expected.port, host, port)
			}
		})
	}
}
