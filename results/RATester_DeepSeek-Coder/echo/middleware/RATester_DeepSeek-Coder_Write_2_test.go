package middleware

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Test case 1",
			input:    []byte("Hello, world!"),
			expected: 13,
		},
		{
			name:     "Test case 2",
			input:    []byte(""),
			expected: 0,
		},
		{
			name:     "Test case 3",
			input:    nil,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			writer := &bodyDumpResponseWriter{
				Writer: &bytes.Buffer{},
			}

			n, err := writer.Write(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if n != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, n)
			}
		})
	}
}
