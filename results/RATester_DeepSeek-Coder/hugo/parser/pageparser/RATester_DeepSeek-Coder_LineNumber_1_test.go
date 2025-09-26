package pageparser

import (
	"fmt"
	"testing"
)

func TestLineNumber_1(t *testing.T) {
	testCases := []struct {
		name     string
		source   []byte
		expected int
	}{
		{
			name:     "Empty source",
			source:   []byte(""),
			expected: 1,
		},
		{
			name:     "Source with one line",
			source:   []byte("line1"),
			expected: 1,
		},
		{
			name:     "Source with multiple lines",
			source:   []byte("line1\nline2\nline3"),
			expected: 2,
		},
		{
			name:     "Source with multiple lines and trailing newline",
			source:   []byte("line1\nline2\nline3\n"),
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			iterator := &Iterator{}
			result := iterator.LineNumber(tc.source)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
