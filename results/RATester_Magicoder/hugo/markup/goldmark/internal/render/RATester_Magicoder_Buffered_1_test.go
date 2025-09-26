package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffered_1(t *testing.T) {
	b := &BufWriter{
		Buffer: &bytes.Buffer{},
	}

	testCases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "empty buffer",
			input:    []byte{},
			expected: 0,
		},
		{
			name:     "non-empty buffer",
			input:    []byte("hello"),
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b.Buffer.Write(tc.input)
			result := b.Buffered()
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
