package css

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/css"
)

func TestQuoted_1(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name     string
		input    any
		expected css.QuotedString
	}{
		{
			name:     "string input",
			input:    "test",
			expected: css.QuotedString("test"),
		},
		{
			name:     "int input",
			input:    123,
			expected: css.QuotedString("123"),
		},
		{
			name:     "float input",
			input:    123.456,
			expected: css.QuotedString("123.456"),
		},
		{
			name:     "bool input",
			input:    true,
			expected: css.QuotedString("true"),
		},
		{
			name:     "nil input",
			input:    nil,
			expected: css.QuotedString(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.Quoted(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
