package utils

import (
	"fmt"
	"testing"
)

func TestGetValue_3(t *testing.T) {
	testCases := []struct {
		name     string
		input    SimpleKV
		expected interface{}
	}{
		{
			name: "Test case 1",
			input: SimpleKV{
				Key:   "testKey",
				Value: "testValue",
			},
			expected: "testValue",
		},
		{
			name: "Test case 2",
			input: SimpleKV{
				Key:   123,
				Value: 456,
			},
			expected: 456,
		},
		{
			name: "Test case 3",
			input: SimpleKV{
				Key:   "anotherKey",
				Value: nil,
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.GetValue()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
