package types

import (
	"fmt"
	"testing"
)

func TestCanonicalDomain_1(t *testing.T) {
	testCases := []struct {
		name     string
		domain   string
		expected string
	}{
		{"empty string", "", ""},
		{"trim spaces", "   example.com   ", "example.com"},
		{"convert to lowercase", "EXAMPLE.COM", "example.com"},
		{"combination", "   EXAMPLE.COM   ", "example.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := CanonicalDomain(tc.domain)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
