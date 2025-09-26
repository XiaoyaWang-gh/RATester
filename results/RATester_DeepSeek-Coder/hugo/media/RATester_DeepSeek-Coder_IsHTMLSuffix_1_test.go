package media

import (
	"fmt"
	"testing"
)

func TestIsHTMLSuffix_1(t *testing.T) {
	tests := []struct {
		name     string
		suffix   string
		expected bool
	}{
		{
			name:     "HTML suffix",
			suffix:   ".html",
			expected: true,
		},
		{
			name:     "Markdown suffix",
			suffix:   ".md",
			expected: false,
		},
		{
			name:     "AsciiDoc suffix",
			suffix:   ".adoc",
			expected: false,
		},
		{
			name:     "Pandoc suffix",
			suffix:   ".md",
			expected: false,
		},
		{
			name:     "ReStructuredText suffix",
			suffix:   ".rst",
			expected: false,
		},
		{
			name:     "EmacsOrgMode suffix",
			suffix:   ".org",
			expected: false,
		},
		{
			name:     "Unknown suffix",
			suffix:   ".unknown",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			contentTypes := ContentTypes{
				HTML: Type{
					MainType:    "text",
					SubType:     "html",
					SuffixesCSV: ".html,.htm",
				},
				Markdown: Type{
					MainType:    "text",
					SubType:     "markdown",
					SuffixesCSV: ".md,.markdown",
				},
				// Add other types as needed...
			}

			result := contentTypes.IsHTMLSuffix(test.suffix)

			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
