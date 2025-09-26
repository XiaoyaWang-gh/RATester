package bytesconv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToBytes_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []byte{},
		},
		{
			name:     "Single character",
			input:    "a",
			expected: []byte{'a'},
		},
		{
			name:     "Multiple characters",
			input:    "hello",
			expected: []byte{'h', 'e', 'l', 'l', 'o'},
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

			result := StringToBytes(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
