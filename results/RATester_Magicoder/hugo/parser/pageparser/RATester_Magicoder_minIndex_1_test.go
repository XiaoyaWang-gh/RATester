package pageparser

import (
	"fmt"
	"testing"
)

func TestminIndex_1(t *testing.T) {
	tests := []struct {
		name     string
		indices  []int
		expected int
	}{
		{
			name:     "Test case 1",
			indices:  []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Test case 2",
			indices:  []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		{
			name:     "Test case 3",
			indices:  []int{1, -2, 3, -4, 5},
			expected: 1,
		},
		{
			name:     "Test case 4",
			indices:  []int{},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := minIndex(test.indices...)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
