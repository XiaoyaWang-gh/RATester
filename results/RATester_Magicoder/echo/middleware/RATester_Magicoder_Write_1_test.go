package middleware

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	testCases := []struct {
		name          string
		ignoreWrites  bool
		input         []byte
		expectedBytes int
		expectedError error
	}{
		{
			name:          "Ignore writes",
			ignoreWrites:  true,
			input:         []byte("test"),
			expectedBytes: 4,
			expectedError: nil,
		},
		{
			name:          "Do not ignore writes",
			ignoreWrites:  false,
			input:         []byte("test"),
			expectedBytes: 4,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &ignorableWriter{
				ignoreWrites: tc.ignoreWrites,
			}
			n, err := w.Write(tc.input)
			if n != tc.expectedBytes {
				t.Errorf("Expected %d bytes, got %d", tc.expectedBytes, n)
			}
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
