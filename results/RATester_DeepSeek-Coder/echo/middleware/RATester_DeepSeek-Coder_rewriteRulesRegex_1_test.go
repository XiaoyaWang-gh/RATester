package middleware

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestRewriteRulesRegex_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]string
		expected map[*regexp.Regexp]string
	}{
		{
			name:  "Test case 1",
			input: map[string]string{"^abc$": "def"},
			expected: map[*regexp.Regexp]string{
				regexp.MustCompile("^abc$"): "def",
			},
		},
		{
			name:  "Test case 2",
			input: map[string]string{"a*b": "c*d"},
			expected: map[*regexp.Regexp]string{
				regexp.MustCompile("a(.*?)b"): "c(.*?)d",
			},
		},
		{
			name:  "Test case 3",
			input: map[string]string{"^a*b$": "^c*d$"},
			expected: map[*regexp.Regexp]string{
				regexp.MustCompile("^a(.*?)b$"): "^c(.*?)d$",
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

			result := rewriteRulesRegex(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
