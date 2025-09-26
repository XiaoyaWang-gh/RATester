package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCheckAllowedHTTPMethod_1(t *testing.T) {
	type testCase struct {
		name     string
		method   string
		expected error
	}

	testCases := []testCase{
		{
			name:     "Allowed HTTP Method",
			method:   "GET",
			expected: nil,
		},
		{
			name:   "Disallowed HTTP Method",
			method: "POST",
			expected: &AccessDeniedError{
				name:     "POST",
				path:     "security.http.method",
				policies: "<policies>",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			config := Config{
				HTTP: HTTP{
					Methods: Whitelist{
						acceptNone: false,
						patterns: []*regexp.Regexp{
							regexp.MustCompile("GET"),
						},
						patternsStrings: []string{"GET"},
					},
				},
			}

			err := config.CheckAllowedHTTPMethod(tc.method)
			if err != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, err)
			}
		})
	}
}
