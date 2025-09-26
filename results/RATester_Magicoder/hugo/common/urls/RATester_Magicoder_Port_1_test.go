package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPort_1(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  BaseURL
		expected int
	}{
		{
			name: "Port 80",
			baseURL: BaseURL{
				url: &url.URL{
					Host: "localhost:80",
				},
			},
			expected: 80,
		},
		{
			name: "Port 443",
			baseURL: BaseURL{
				url: &url.URL{
					Host: "localhost:443",
				},
			},
			expected: 443,
		},
		{
			name: "No Port",
			baseURL: BaseURL{
				url: &url.URL{
					Host: "localhost",
				},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := test.baseURL.Port()
			if got != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, got)
			}
		})
	}
}
