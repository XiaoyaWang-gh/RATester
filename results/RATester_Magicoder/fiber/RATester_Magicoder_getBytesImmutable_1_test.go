package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetBytesImmutable_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "Test Case 1",
			input:    "Hello, World!",
			expected: []byte("Hello, World!"),
		},
		{
			name:     "Test Case 2",
			input:    "1234567890",
			expected: []byte("1234567890"),
		},
		{
			name:     "Test Case 3",
			input:    "",
			expected: []byte(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := getBytesImmutable(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
