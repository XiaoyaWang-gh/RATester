package http

import (
	"fmt"
	"testing"
)

func TestCanonicalHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"example.com", "example.com"},
		{"EXAMPLE.COM", "example.com"},
		{"example.com.", "example.com"},
		{"example.com:80", "example.com"},
		{"example.com:8080", "example.com:8080"},
	}

	for _, test := range tests {
		actual, err := CanonicalHost(test.input)
		if err != nil {
			t.Errorf("CanonicalHost(%q) returned error: %v", test.input, err)
			continue
		}
		if actual != test.expected {
			t.Errorf("CanonicalHost(%q) = %q; want %q", test.input, actual, test.expected)
		}
	}
}
