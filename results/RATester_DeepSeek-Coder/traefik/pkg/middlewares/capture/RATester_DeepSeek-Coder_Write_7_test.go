package capture

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWrite_7(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Test with empty input",
			input:    []byte{},
			expected: 0,
		},
		{
			name:     "Test with non-empty input",
			input:    []byte("Hello, World!"),
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := httptest.NewRecorder()
			crw := &captureResponseWriter{rw: rw}

			actual, err := crw.Write(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if actual != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, actual)
			}

			if crw.size != int64(tc.expected) {
				t.Errorf("Expected size %d, but got %d", tc.expected, crw.size)
			}
		})
	}
}
