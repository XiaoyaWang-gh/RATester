package stripprefix

import (
	"fmt"
	"testing"
)

func TestGetPrefixStripped_1(t *testing.T) {
	testCases := []struct {
		name     string
		urlPath  string
		prefix   string
		expected string
	}{
		{
			name:     "Test with empty prefix",
			urlPath:  "/api/v1/test",
			prefix:   "",
			expected: "/api/v1/test",
		},
		{
			name:     "Test with non-empty prefix",
			urlPath:  "/api/v1/test",
			prefix:   "/api",
			expected: "/v1/test",
		},
		{
			name:     "Test with non-empty prefix and trailing slash",
			urlPath:  "/api/v1/test/",
			prefix:   "/api",
			expected: "/v1/test/",
		},
		{
			name:     "Test with non-empty prefix and multiple trailing slashes",
			urlPath:  "/api/v1/test///",
			prefix:   "/api",
			expected: "/v1/test///",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &stripPrefix{}
			result := s.getPrefixStripped(tc.urlPath, tc.prefix)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
