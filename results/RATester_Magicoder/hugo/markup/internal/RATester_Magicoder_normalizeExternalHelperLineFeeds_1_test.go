package internal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestnormalizeExternalHelperLineFeeds_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected []byte
	}{
		{
			name:     "Test with no line feeds",
			content:  []byte("test"),
			expected: []byte("test"),
		},
		{
			name:     "Test with line feeds",
			content:  []byte("test\r\n"),
			expected: []byte("test\n"),
		},
		{
			name:     "Test with multiple line feeds",
			content:  []byte("test\r\n\r\n"),
			expected: []byte("test\n\n"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizeExternalHelperLineFeeds(test.content)
			if !bytes.Equal(result, test.expected) {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
