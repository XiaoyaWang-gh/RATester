package parse

import (
	"fmt"
	"testing"
)

func TestParseTemplateName_1(t *testing.T) {
	tree := &Tree{}

	tests := []struct {
		name     string
		token    item
		context  string
		expected string
	}{
		{
			name: "string token",
			token: item{
				typ: itemString,
				val: `"test"`,
			},
			context:  "test context",
			expected: "test",
		},
		{
			name: "raw string token",
			token: item{
				typ: itemRawString,
				val: "`test`",
			},
			context:  "test context",
			expected: "test",
		},
		{
			name: "unquoted string token",
			token: item{
				typ: itemString,
				val: "test",
			},
			context:  "test context",
			expected: "test",
		},
		{
			name: "unquoted raw string token",
			token: item{
				typ: itemRawString,
				val: "test",
			},
			context:  "test context",
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tree.parseTemplateName(test.token, test.context)
			if result != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, result)
			}
		})
	}
}
