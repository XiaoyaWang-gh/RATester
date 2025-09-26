package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAccept_2(t *testing.T) {
	type testCase struct {
		name      string
		whitelist Whitelist
		input     string
		expected  bool
	}

	testCases := []testCase{
		{
			name: "AcceptNoneTrue",
			whitelist: Whitelist{
				acceptNone: true,
			},
			input:    "test",
			expected: false,
		},
		{
			name: "MatchFound",
			whitelist: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("test"),
				},
			},
			input:    "test",
			expected: true,
		},
		{
			name: "MatchNotFound",
			whitelist: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("notfound"),
				},
			},
			input:    "test",
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

			result := tc.whitelist.Accept(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
