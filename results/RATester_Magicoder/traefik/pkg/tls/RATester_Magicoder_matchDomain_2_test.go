package tls

import (
	"fmt"
	"testing"
)

func TestMatchDomain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		serverName string
		certDomain string
		expected   bool
	}{
		{"example.com", "example.com", true},
		{"www.example.com", "example.com", true},
		{"example.com", "*.example.com", true},
		{"www.example.com", "*.example.com", true},
		{"example.com", "example.com.", true},
		{"www.example.com", "example.com.", true},
		{"example.com", "*.example.com.", true},
		{"www.example.com", "*.example.com.", true},
		{"example.com", "com", false},
		{"www.example.com", "com", false},
		{"example.com", "*.com", false},
		{"www.example.com", "*.com", false},
		{"example.com", "example.com.", false},
		{"www.example.com", "example.com.", false},
		{"example.com", "*.example.com.", false},
		{"www.example.com", "*.example.com.", false},
	}

	for _, test := range tests {
		result := matchDomain(test.serverName, test.certDomain)
		if result != test.expected {
			t.Errorf("matchDomain(%s, %s) = %t; want %t", test.serverName, test.certDomain, result, test.expected)
		}
	}
}
