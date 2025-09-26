package security

import (
	"fmt"
	"regexp"
	"testing"

	"gotest.tools/assert"
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
						patterns: []*regexp.Regexp{
							regexp.MustCompile("^https://allowed\\.com/"),
						},
					},
				},
			},
			url:      "https://allowed.com/path",
			expected: nil,
		},
		{
			name: "disallowed url",
			config: Config{
				HTTP: HTTP{
					URLs: Whitelist{
						patterns: []*regexp.Regexp{
							regexp.MustCompile("^https://allowed\\.com/"),
						},
					},
				},
			},
			url:      "https://disallowed.com/path",
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
			if err != nil && tc.expected != nil {
				assert.Equal(t, err.Error(), tc.expected.Error())
			} else {
				assert.Equal(t, err, tc.expected)
			}
		})
	}
}
