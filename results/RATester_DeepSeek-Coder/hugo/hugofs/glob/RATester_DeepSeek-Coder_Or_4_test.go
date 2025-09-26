package glob

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestOr_4(t *testing.T) {
	type testCase struct {
		name     string
		globs    []glob.Glob
		input    string
		expected bool
	}

	testCases := []testCase{
		{
			name: "match first",
			globs: []glob.Glob{
				glob.MustCompile("a*"),
				glob.MustCompile("b*"),
			},
			input:    "abc",
			expected: true,
		},
		{
			name: "match second",
			globs: []glob.Glob{
				glob.MustCompile("c*"),
				glob.MustCompile("d*"),
			},
			input:    "def",
			expected: true,
		},
		{
			name: "no match",
			globs: []glob.Glob{
				glob.MustCompile("e*"),
				glob.MustCompile("f*"),
			},
			input:    "ghi",
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

			result := Or(tc.globs...).Match(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
