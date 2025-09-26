package requestdecorator

import (
	"fmt"
	"testing"
)

func TestParseHost_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"IPv4 without port", "192.168.1.1", "192.168.1.1"},
		{"IPv4 with port", "192.168.1.1:8080", "192.168.1.1"},
		{"IPv6 without port", "[2001:0db8:85a3:0000:0000:8a2e:0370:7334]", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{"IPv6 with port", "[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:8080", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{"Invalid address", "invalid", "invalid"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := parseHost(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
