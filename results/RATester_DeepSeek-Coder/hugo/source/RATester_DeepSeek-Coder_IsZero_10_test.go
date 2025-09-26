package source

import (
	"fmt"
	"testing"
)

func TestIsZero_10(t *testing.T) {
	type testCase struct {
		name     string
		input    GitInfo
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test with empty GitInfo",
			input: GitInfo{
				Hash: "",
			},
			expected: true,
		},
		{
			name: "Test with non-empty GitInfo",
			input: GitInfo{
				Hash: "abc123",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.IsZero()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
