package constants

import (
	"fmt"
	"testing"
)

func TestIsResourceTransformationPermalinkHash_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Test case 1",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: false,
		},
		{
			name:     "Test case 2",
			input:    "0123456789",
			expected: false,
		},
		{
			name:     "Test case 3",
			input:    "abcdefghijklmnopqrstuvwxyz0123456789",
			expected: false,
		},
		{
			name:     "Test case 4",
			input:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			expected: false,
		},
		{
			name:     "Test case 5",
			input:    ResourceTransformationFingerprint,
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

			result := IsResourceTransformationPermalinkHash(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
