package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCheckAllowedExec_1(t *testing.T) {
	type testCase struct {
		name     string
		config   Config
		input    string
		expected error
	}

	testCases := []testCase{
		{
			name: "Allowed Exec",
			config: Config{
				Exec: Exec{
					Allow: Whitelist{
						patterns: []*regexp.Regexp{regexp.MustCompile("^/bin/echo$")},
					},
				},
			},
			input:    "/bin/echo",
			expected: nil,
		},
		{
			name: "Not Allowed Exec",
			config: Config{
				Exec: Exec{
					Allow: Whitelist{
						patterns: []*regexp.Regexp{regexp.MustCompile("^/bin/echo$")},
					},
				},
			},
			input:    "/bin/sh",
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

			err := tc.config.CheckAllowedExec(tc.input)
			if err != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, err)
			}
		})
	}
}
