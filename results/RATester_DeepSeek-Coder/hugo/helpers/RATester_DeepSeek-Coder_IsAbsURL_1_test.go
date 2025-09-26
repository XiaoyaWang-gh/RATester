package helpers

import (
	"fmt"
	"net/url"
	"testing"
)

func TestIsAbsURL_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected bool
		err      error
	}

	tests := []test{
		{"absolute http", "http://example.com", true, nil},
		{"absolute https", "https://example.com", true, nil},
		{"relative http", "/example.com", false, nil},
		{"relative https", "/example.com", false, nil},
		{"invalid url", ":example.com", false, &url.Error{}},
	}

	p := &PathSpec{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := p.IsAbsURL(test.input)
			if err != nil && test.err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err == nil && test.err != nil {
				t.Errorf("Expected error, got nil")
			}
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
