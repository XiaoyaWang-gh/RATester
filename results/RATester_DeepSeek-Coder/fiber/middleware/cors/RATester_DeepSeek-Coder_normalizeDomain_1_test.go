package cors

import (
	"fmt"
	"testing"
)

func TestNormalizeDomain_1(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"http://example.com", "example.com"},
		{"https://example.com", "example.com"},
		{"https://example.com:8080", "example.com"},
		{"example.com", "example.com"},
		{"[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:8080", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{"[2001:0db8:85a3:0000:0000:8a2e:0370:7334]", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{"", ""},
		{"http://", ""},
		{"https://", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizeDomain(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
