package internal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNormalizeExternalHelperLineFeeds_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Test with CRLF line endings",
			input:    []byte("line1\r\nline2\r\nline3"),
			expected: []byte("line1\nline2\nline3"),
		},
		{
			name:     "Test with LF line endings",
			input:    []byte("line1\nline2\nline3"),
			expected: []byte("line1\nline2\nline3"),
		},
		{
			name:     "Test with CR line endings",
			input:    []byte("line1\rline2\rline3"),
			expected: []byte("line1\nline2\nline3"),
		},
		{
			name:     "Test with no line endings",
			input:    []byte("line1line2line3"),
			expected: []byte("line1line2line3"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizeExternalHelperLineFeeds(tc.input)
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
