package headers

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestIsOriginAllowed_1(t *testing.T) {
	type testCase struct {
		name     string
		origin   string
		expected bool
		result   string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			origin:   "http://example.com",
			expected: true,
			result:   "http://example.com",
		},
		{
			name:     "Test case 2",
			origin:   "http://foo.example.com",
			expected: true,
			result:   "http://foo.example.com",
		},
		{
			name:     "Test case 3",
			origin:   "http://bar.example.com",
			expected: false,
			result:   "",
		},
		{
			name:     "Test case 4",
			origin:   "*",
			expected: true,
			result:   "*",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := &Header{
				headers: &dynamic.Headers{
					AccessControlAllowOriginList: []string{"http://example.com", "http://foo.example.com", "*"},
				},
				allowOriginRegexes: []*regexp.Regexp{regexp.MustCompile("^http://bar.example.com$")},
			}

			actual, result := h.isOriginAllowed(test.origin)
			if actual != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
			if result != test.result {
				t.Errorf("Expected %v, but got %v", test.result, result)
			}
		})
	}
}
