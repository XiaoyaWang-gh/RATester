package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStripEmptyNav_1(t *testing.T) {
	type test struct {
		name     string
		input    []byte
		expected []byte
	}

	tests := []test{
		{
			name:     "Test case 1",
			input:    []byte("<nav>\n</nav>\n\n"),
			expected: []byte(""),
		},
		{
			name:     "Test case 2",
			input:    []byte("<nav>\n</nav>\n\nHello World"),
			expected: []byte("Hello World"),
		},
		{
			name:     "Test case 3",
			input:    []byte("Hello World\n<nav>\n</nav>\n\n"),
			expected: []byte("Hello World"),
		},
		{
			name:     "Test case 4",
			input:    []byte("<nav>\n</nav>\n\n<nav>\n</nav>\n\n"),
			expected: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := stripEmptyNav(tt.input)
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}
