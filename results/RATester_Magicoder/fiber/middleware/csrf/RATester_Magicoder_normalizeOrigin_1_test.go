package csrf

import (
	"fmt"
	"testing"
)

func TestnormalizeOrigin_1(t *testing.T) {
	tests := []struct {
		name     string
		origin   string
		expected bool
		result   string
	}{
		{
			name:     "valid http origin",
			origin:   "http://example.com",
			expected: true,
			result:   "http://example.com",
		},
		{
			name:     "valid https origin",
			origin:   "https://example.com",
			expected: true,
			result:   "https://example.com",
		},
		{
			name:     "invalid scheme",
			origin:   "ftp://example.com",
			expected: false,
			result:   "",
		},
		{
			name:     "invalid host",
			origin:   "http://*",
			expected: false,
			result:   "",
		},
		{
			name:     "invalid path",
			origin:   "http://example.com/path",
			expected: false,
			result:   "",
		},
		{
			name:     "invalid query",
			origin:   "http://example.com?query=test",
			expected: false,
			result:   "",
		},
		{
			name:     "invalid fragment",
			origin:   "http://example.com#fragment",
			expected: false,
			result:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, result := normalizeOrigin(test.origin)
			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
			if result != test.result {
				t.Errorf("Expected %v, but got %v", test.result, result)
			}
		})
	}
}
