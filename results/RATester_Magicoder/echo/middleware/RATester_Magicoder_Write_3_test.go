package middleware

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_3(t *testing.T) {
	w := &gzipResponseWriter{
		wroteHeader:       false,
		wroteBody:         false,
		minLength:         100,
		minLengthExceeded: false,
		buffer:            bytes.NewBuffer([]byte{}),
		code:              200,
	}

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
			input:    []byte("This is a longer test case."),
			expected: 26,
		},
		{
			name:     "Test case 3",
			input:    []byte("This is a very long test case. It should exceed the minimum length."),
			expected: 60,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			n, err := w.Write(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if n != tc.expected {
				t.Errorf("Expected %d bytes, got %d", tc.expected, n)
			}
		})
	}
}
