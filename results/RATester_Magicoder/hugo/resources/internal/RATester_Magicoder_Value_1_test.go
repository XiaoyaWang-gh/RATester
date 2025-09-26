package internal

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	type testCase struct {
		name     string
		elements []any
		expected string
	}

	testCases := []testCase{
		{
			name:     "test1",
			elements: []any{1, 2, 3},
			expected: "test1_123",
		},
		{
			name:     "test2",
			elements: []any{},
			expected: "test2",
		},
		{
			name:     "test3",
			elements: []any{4, 5, 6},
			expected: "test3_456",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			key := ResourceTransformationKey{
				Name:     tc.name,
				elements: tc.elements,
			}

			result := key.Value()

			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
