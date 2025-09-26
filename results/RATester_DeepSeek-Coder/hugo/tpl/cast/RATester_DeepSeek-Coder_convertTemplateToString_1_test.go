package cast

import (
	"fmt"
	"html/template"
	"testing"
)

func TestConvertTemplateToString_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    any
		expected any
	}

	testCases := []testCase{
		{
			name:     "Test with template.HTML",
			input:    template.HTML("<p>Hello, World</p>"),
			expected: "<p>Hello, World</p>",
		},
		{
			name:     "Test with template.CSS",
			input:    template.CSS("body { color: red; }"),
			expected: "body { color: red; }",
		},
		{
			name:     "Test with template.HTMLAttr",
			input:    template.HTMLAttr("onclick=\"doSomething()\""),
			expected: "onclick=\"doSomething()\"",
		},
		{
			name:     "Test with template.JS",
			input:    template.JS("console.log('Hello, World')"),
			expected: "console.log('Hello, World')",
		},
		{
			name:     "Test with template.JSStr",
			input:    template.JSStr("'Hello, World'"),
			expected: "'Hello, World'",
		},
		{
			name:     "Test with string",
			input:    "Hello, World",
			expected: "Hello, World",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := convertTemplateToString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
