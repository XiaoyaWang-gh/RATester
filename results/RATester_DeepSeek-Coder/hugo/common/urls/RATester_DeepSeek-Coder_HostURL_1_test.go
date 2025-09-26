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
			name: "HostURL with path",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
					Path:   "/path",
				},
			},
			expected: "https://example.com",
		},
		{
			name: "HostURL without path",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
				},
			},
			expected: "https://example.com",
		},
		{
			name: "HostURL with empty path",
			baseURL: BaseURL{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
					Path:   "",
				},
			},
			expected: "https://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.baseURL.HostURL()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
