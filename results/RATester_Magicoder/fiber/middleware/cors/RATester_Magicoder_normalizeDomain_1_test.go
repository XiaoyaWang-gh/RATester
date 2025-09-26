package cors

import (
	"fmt"
	"testing"
)

func TestnormalizeDomain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		input    string
		expected string
	}{
		{"http://www.example.com", "www.example.com"},
		{"https://www.example.com", "www.example.com"},
		{"www.example.com", "www.example.com"},
		{"http://www.example.com:8080", "www.example.com"},
		{"https://www.example.com:443", "www.example.com"},
		{"www.example.com:8080", "www.example.com"},
		{"[::1]:8080", "::1"},
		{"[2001:db8::1]:8080", "2001:db8::1"},
		{"http://[::1]", "::1"},
		{"https://[2001:db8::1]", "2001:db8::1"},
		{"[::1]", "::1"},
		{"[2001:db8::1]", "2001:db8::1"},
	}

	for _, test := range tests {
		result := normalizeDomain(test.input)
		if result != test.expected {
			t.Errorf("normalizeDomain(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
