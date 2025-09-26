package helpers

import (
	"fmt"
	"testing"
)

func TestPermalinkForBaseURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		link     string
		baseURL  string
		expected string
	}

	tests := []test{
		{"/test", "https://example.com", "https://example.com/test"},
		{"/test", "https://example.com/", "https://example.com/test"},
		{"/test/", "https://example.com", "https://example.com/test/"},
		{"/test/", "https://example.com/", "https://example.com/test/"},
		{"test", "https://example.com", "https://example.com/test"},
		{"test/", "https://example.com", "https://example.com/test/"},
	}

	for _, test := range tests {
		p := &PathSpec{}
		result := p.PermalinkForBaseURL(test.link, test.baseURL)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
