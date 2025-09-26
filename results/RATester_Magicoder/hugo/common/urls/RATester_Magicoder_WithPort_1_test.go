package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestWithPort_1(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  BaseURL
		port     int
		expected string
	}{
		{
			name:     "WithPort",
			baseURL:  BaseURL{url: &url.URL{Scheme: "http", Host: "example.com"}},
			port:     8080,
			expected: "http://example.com:8080",
		},
		{
			name:     "WithPort",
			baseURL:  BaseURL{url: &url.URL{Scheme: "https", Host: "example.com"}},
			port:     443,
			expected: "https://example.com:443",
		},
		{
			name:     "WithPort",
			baseURL:  BaseURL{url: &url.URL{Scheme: "mailto", Host: "example.com"}},
			port:     25,
			expected: "mailto://example.com:25",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := test.baseURL.WithPort(test.port)
			if err != nil {
				t.Errorf("WithPort() error = %v", err)
				return
			}
			if result.String() != test.expected {
				t.Errorf("WithPort() = %v, want %v", result.String(), test.expected)
			}
		})
	}
}
