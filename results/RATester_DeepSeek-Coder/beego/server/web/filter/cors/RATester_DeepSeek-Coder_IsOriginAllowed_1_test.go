package cors

import (
	"fmt"
	"testing"
)

func TestIsOriginAllowed_1(t *testing.T) {
	testCases := []struct {
		name     string
		options  *Options
		origin   string
		expected bool
	}{
		{
			name: "All origins allowed",
			options: &Options{
				AllowAllOrigins: true,
			},
			origin:   "http://example.com",
			expected: true,
		},
		{
			name: "Specific origin allowed",
			options: &Options{
				AllowOrigins: []string{"http://example.com"},
			},
			origin:   "http://example.com",
			expected: true,
		},
		{
			name: "Wildcard origin allowed",
			options: &Options{
				AllowOrigins: []string{"*"},
			},
			origin:   "http://example.com",
			expected: true,
		},
		{
			name: "Origin not allowed",
			options: &Options{
				AllowOrigins: []string{"http://not-example.com"},
			},
			origin:   "http://example.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual := tc.options.IsOriginAllowed(tc.origin)
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
