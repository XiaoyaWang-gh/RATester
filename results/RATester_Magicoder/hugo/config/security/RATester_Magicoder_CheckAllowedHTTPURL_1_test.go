package security

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestCheckAllowedHTTPURL_1(t *testing.T) {
	type testCase struct {
		name     string
		config   Config
		url      string
		expected error
	}

	testCases := []testCase{
		{
			name: "allowed url",
			config: Config{
				HTTP: HTTP{
					URLs: Whitelist{
						patterns: []*regexp.Regexp{regexp.MustCompile("^https://allowed\\.com$")},
					},
				},
			},
			url:      "https://allowed.com",
			expected: nil,
		},
		{
			name: "disallowed url",
			config: Config{
				HTTP: HTTP{
					URLs: Whitelist{
						patterns: []*regexp.Regexp{regexp.MustCompile("^https://allowed\\.com$")},
					},
				},
			},
			url:      "https://disallowed.com",
			expected: &AccessDeniedError{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.config.CheckAllowedHTTPURL(tc.url)
			if err != nil && tc.expected == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if err == nil && tc.expected != nil {
				t.Errorf("Expected error, but got nil")
			}
			if err != nil && tc.expected != nil {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.expected) {
					t.Errorf("Expected error of type %T, but got %T", tc.expected, err)
				}
			}
		})
	}
}
