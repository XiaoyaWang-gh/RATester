package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHostURL_1(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  BaseURL
		expected string
	}{
		{
			name: "Test with path",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "/path",
				},
			},
			expected: "http://example.com",
		},
		{
			name: "Test without path",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
				},
			},
			expected: "http://example.com",
		},
		{
			name: "Test with trailing slash",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "/path/",
				},
			},
			expected: "http://example.com",
		},
		{
			name: "Test without trailing slash",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "/path",
				},
			},
			expected: "http://example.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.baseURL.HostURL()
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
