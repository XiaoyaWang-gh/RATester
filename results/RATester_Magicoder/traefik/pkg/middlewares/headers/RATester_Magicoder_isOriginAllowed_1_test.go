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
			name:     "Origin is allowed",
			origin:   "example.com",
			expected: true,
			result:   "example.com",
		},
		{
			name:     "Origin is not allowed",
			origin:   "notallowed.com",
			expected: false,
			result:   "",
		},
		{
			name:     "Origin is allowed with wildcard",
			origin:   "wildcard.com",
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
					AccessControlAllowOriginList: []string{"example.com", "*"},
				},
				allowOriginRegexes: []*regexp.Regexp{regexp.MustCompile("wildcard.com")},
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
