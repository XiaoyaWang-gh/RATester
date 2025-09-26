package metrics

import (
	"fmt"
	"testing"
)

func TestParseDatadogAddress_1(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected struct {
			network string
			addr    string
		}
	}{
		{
			name:    "empty address",
			address: "",
			expected: struct {
				network string
				addr    string
			}{
				network: "udp",
				addr:    "localhost:8125",
			},
		},
		{
			name:    "unix address",
			address: unixAddressPrefix + "unix:///var/run/datadog/dsd.socket",
			expected: struct {
				network string
				addr    string
			}{
				network: "unix",
				addr:    "unix:///var/run/datadog/dsd.socket",
			},
		},
		{
			name:    "udp address",
			address: "192.168.0.1:8125",
			expected: struct {
				network string
				addr    string
			}{
				network: "udp",
				addr:    "192.168.0.1:8125",
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

			network, addr := parseDatadogAddress(test.address)
			if network != test.expected.network {
				t.Errorf("Expected network '%s', got '%s'", test.expected.network, network)
			}
			if addr != test.expected.addr {
				t.Errorf("Expected addr '%s', got '%s'", test.expected.addr, addr)
			}
		})
	}
}
