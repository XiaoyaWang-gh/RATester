package bundler

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
)

func TestRead_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Test case 1",
			input:    []byte("test input"),
			expected: []byte("test input"),
		},
		{
			name:     "Test case 2",
			input:    []byte("another test input"),
			expected: []byte("another test input"),
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

			r := &multiReadSeekCloser{
				mr:      bytes.NewReader(tc.input),
				sources: []hugio.ReadSeekCloser{},
			}

			output := make([]byte, len(tc.input))
			n, err := r.Read(output)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if n != len(tc.input) {
				t.Errorf("Expected to read %d bytes, but read %d", len(tc.input), n)
			}

			if !bytes.Equal(output, tc.expected) {
				t.Errorf("Expected output to be %v, but was %v", tc.expected, output)
			}
		})
	}
}
