package tableofcontents

import (
	"fmt"
	"strings"
	"testing"
)

func TestIndent_1(t *testing.T) {
	type testCase struct {
		name     string
		b        *tocBuilder
		n        int
		expected string
	}

	testCases := []testCase{
		{
			name: "Test indent with 1 level",
			b: &tocBuilder{
				s: strings.Builder{},
			},
			n:        1,
			expected: "  ",
		},
		{
			name: "Test indent with 2 levels",
			b: &tocBuilder{
				s: strings.Builder{},
			},
			n:        2,
			expected: "    ",
		},
		{
			name: "Test indent with 0 levels",
			b: &tocBuilder{
				s: strings.Builder{},
			},
			n:        0,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.b.indent(tc.n)
			result := tc.b.s.String()
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
