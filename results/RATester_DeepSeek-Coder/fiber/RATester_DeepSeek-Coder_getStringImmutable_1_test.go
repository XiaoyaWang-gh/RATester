package fiber

import (
	"fmt"
	"testing"
)

func TestGetStringImmutable_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Test case 1",
			input:    []byte("test"),
			expected: "test",
		},
		{
			name:     "Test case 2",
			input:    []byte("hello"),
			expected: "hello",
		},
		{
			name:     "Test case 3",
			input:    []byte("world"),
			expected: "world",
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
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
