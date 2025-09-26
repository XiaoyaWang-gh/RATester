package middleware

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_3(t *testing.T) {
	testCases := []struct {
		name           string
		minLength      int
		input          []byte
		expectedOutput []byte
		expectedError  error
	}{
		{
			name:           "Test case 1",
			minLength:      10,
			input:          []byte("Hello, world!"),
			expectedOutput: []byte("Hello, world!"),
			expectedError:  nil,
		},
		{
			name:           "Test case 2",
			minLength:      10,
			input:          []byte("Hello, world!"),
			expectedOutput: []byte("Hello, world!"),
			expectedError:  nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			// Create a new gzipResponseWriter
			w := &gzipResponseWriter{
				minLength: tc.minLength,
				buffer:    &bytes.Buffer{},
			}

			// Call the Write method
			n, err := w.Write(tc.input)

			// Check the output
			if n != len(tc.input) {
				t.Errorf("Expected %d bytes written, got %d", len(tc.input), n)
			}
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
			if !bytes.Equal(w.buffer.Bytes(), tc.expectedOutput) {
				t.Errorf("Expected output %v, got %v", tc.expectedOutput, w.buffer.Bytes())
			}
		})
	}
}
