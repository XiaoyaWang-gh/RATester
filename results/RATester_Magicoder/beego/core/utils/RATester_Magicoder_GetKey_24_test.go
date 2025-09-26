package utils

import (
	"fmt"
	"testing"
)

func TestGetKey_24(t *testing.T) {
	testCases := []struct {
		name     string
		key      interface{}
		expected interface{}
	}{
		{
			name:     "Test case 1",
			key:      "testKey",
			expected: "testKey",
		},
		{
			name:     "Test case 2",
			key:      123,
			expected: 123,
		},
		{
			name:     "Test case 3",
			key:      true,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &SimpleKV{
				Key: tc.key,
			}

			result := s.GetKey()

			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
