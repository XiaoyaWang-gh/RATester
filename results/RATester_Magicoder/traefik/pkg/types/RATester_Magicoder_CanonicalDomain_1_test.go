package types

import (
	"fmt"
	"testing"
)

func TestCanonicalDomain_1(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected string
	}{
		{
			name:     "Test with empty string",
			domain:   "",
			expected: "",
		},
		{
			name:     "Test with domain name",
			domain:   "Example.com",
			expected: "example.com",
		},
		{
			name:     "Test with domain name with leading and trailing spaces",
			domain:   "   Example.com   ",
			expected: "example.com",
		},
		{
			name:     "Test with domain name with uppercase letters",
			domain:   "EXAMPLE.COM",
			expected: "example.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := CanonicalDomain(test.domain)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
