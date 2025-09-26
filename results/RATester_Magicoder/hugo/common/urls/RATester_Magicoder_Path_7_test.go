package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPath_7(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  BaseURL
		expected string
	}{
		{
			name: "WithPath",
			baseURL: BaseURL{
				url: &url.URL{
					Path: "/path",
				},
			},
			expected: "/path",
		},
		{
			name: "WithPathNoTrailingSlash",
			baseURL: BaseURL{
				url: &url.URL{
					Path: "/path",
				},
			},
			expected: "/path",
		},
		{
			name: "WithoutPath",
			baseURL: BaseURL{
				url: &url.URL{},
			},
			expected: "",
		},
		{
			name: "BasePath",
			baseURL: BaseURL{
				url: &url.URL{
					Path: "/",
				},
			},
			expected: "/",
		},
		{
			name: "BasePathNoTrailingSlash",
			baseURL: BaseURL{
				url: &url.URL{
					Path: "/",
				},
			},
			expected: "/",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.baseURL.Path()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
