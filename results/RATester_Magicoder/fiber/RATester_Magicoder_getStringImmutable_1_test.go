package fiber

import (
	"fmt"
	"testing"
)

func TestgetStringImmutable_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Test case 1",
			input:    []byte("Hello, World!"),
			expected: "Hello, World!",
		},
		{
			name:     "Test case 2",
			input:    []byte(""),
			expected: "",
		},
		{
			name:     "Test case 3",
			input:    []byte("1234567890"),
			expected: "1234567890",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := getStringImmutable(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
