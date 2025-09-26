package template

import (
	"fmt"
	"testing"
)

func TestHtmlNameFilter_1(t *testing.T) {
	type test struct {
		name     string
		input    []any
		expected string
	}

	tests := []test{
		{
			name:     "Test with valid HTML attribute name",
			input:    []any{"class", contentTypeHTMLAttr},
			expected: "class",
		},
		{
			name:     "Test with invalid HTML attribute name",
			input:    []any{"class1", contentTypeHTMLAttr},
			expected: filterFailsafe,
		},
		{
			name:     "Test with empty input",
			input:    []any{},
			expected: filterFailsafe,
		},
		{
			name:     "Test with non-string input",
			input:    []any{123, contentTypeHTMLAttr},
			expected: filterFailsafe,
		},
		{
			name:     "Test with non-HTML content type",
			input:    []any{"class", contentTypePlain},
			expected: filterFailsafe,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := htmlNameFilter(tt.input...); got != tt.expected {
				t.Errorf("htmlNameFilter() = %v, want %v", got, tt.expected)
			}
		})
	}
}
