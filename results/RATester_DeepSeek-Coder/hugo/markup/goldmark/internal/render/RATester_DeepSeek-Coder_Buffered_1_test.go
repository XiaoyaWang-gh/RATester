package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffered_1(t *testing.T) {
	buf := &BufWriter{
		Buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
	}

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty buffer",
			input:    "",
			expected: 0,
		},
		{
			name:     "non-empty buffer",
			input:    "Hello, World",
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			buf.WriteString(tc.input)
			actual := buf.Buffered()
			if actual != tc.expected {
				t.Errorf("Expected Buffered() to return %d, but got %d", tc.expected, actual)
			}
		})
	}
}
