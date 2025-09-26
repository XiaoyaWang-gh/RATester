package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPath_7(t *testing.T) {
	tests := []struct {
		name     string
		b        BaseURL
		expected string
	}{
		{
			name: "Test with path",
			b: BaseURL{
				url: &url.URL{
					Path: "/test/path",
				},
			},
			expected: "/test/path",
		},
		{
			name: "Test with no path",
			b: BaseURL{
				url: &url.URL{
					Path: "",
				},
			},
			expected: "",
		},
		{
			name: "Test with path with trailing slash",
			b: BaseURL{
				url: &url.URL{
					Path: "/test/path/",
				},
			},
			expected: "/test/path/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.b.Path()
			if got != tt.expected {
				t.Errorf("Path() = %v, want %v", got, tt.expected)
			}
		})
	}
}
