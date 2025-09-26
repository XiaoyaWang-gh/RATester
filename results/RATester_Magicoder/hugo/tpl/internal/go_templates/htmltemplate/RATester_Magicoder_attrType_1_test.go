package template

import (
	"fmt"
	"testing"
)

func TestattrType_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected contentType
	}{
		{
			name:     "data- attribute",
			input:    "data-test",
			expected: contentTypePlain,
		},
		{
			name:     "xmlns attribute",
			input:    "xmlns:test",
			expected: contentTypeURL,
		},
		{
			name:     "on attribute",
			input:    "onclick",
			expected: contentTypeJS,
		},
		{
			name:     "src attribute",
			input:    "src",
			expected: contentTypeURL,
		},
		{
			name:     "uri attribute",
			input:    "uri",
			expected: contentTypeURL,
		},
		{
			name:     "url attribute",
			input:    "url",
			expected: contentTypeURL,
		},
		{
			name:     "plain attribute",
			input:    "test",
			expected: contentTypePlain,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := attrType(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
