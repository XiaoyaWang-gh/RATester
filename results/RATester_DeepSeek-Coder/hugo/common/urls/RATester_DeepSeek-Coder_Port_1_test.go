package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPort_1(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected int
	}{
		{
			name:     "Port 80",
			url:      "http://localhost:80",
			expected: 80,
		},
		{
			name:     "Port 443",
			url:      "https://localhost:443",
			expected: 443,
		},
		{
			name:     "No Port",
			url:      "http://localhost",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			u, _ := url.Parse(tt.url)
			b := BaseURL{
				url: u,
			}
			got := b.Port()
			if got != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, got)
			}
		})
	}
}
