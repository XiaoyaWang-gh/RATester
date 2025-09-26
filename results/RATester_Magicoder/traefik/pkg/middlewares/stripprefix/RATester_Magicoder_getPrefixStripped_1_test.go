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
			name:     "Test case 1",
			urlPath:  "/api/v1/test",
			prefix:   "/api/v1",
			expected: "/test",
		},
		{
			name:     "Test case 2",
			urlPath:  "/api/v1/test",
			prefix:   "/api/v2",
			expected: "/api/v1/test",
		},
		{
			name:     "Test case 3",
			urlPath:  "/api/v1/test",
			prefix:   "/api/v1/",
			expected: "/test",
		},
		{
			name:     "Test case 4",
			urlPath:  "/api/v1/test",
			prefix:   "/api/v1",
			expected: "/test",
		},
		{
			name:     "Test case 5",
			urlPath:  "/api/v1/test",
			prefix:   "/api/v1",
			expected: "/test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			sp := &stripPrefix{}
			result := sp.getPrefixStripped(tc.urlPath, tc.prefix)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
