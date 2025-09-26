package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_11(t *testing.T) {
	testCases := []struct {
		name     string
		minSize  MinSize
		input    interface{}
		expected bool
	}{
		{
			name: "string length is greater than or equal to min",
			minSize: MinSize{
				Min: 5,
			},
			input:    "hello",
			expected: true,
		},
		{
			name: "string length is less than min",
			minSize: MinSize{
				Min: 5,
			},
			input:    "hi",
			expected: false,
		},
		{
			name: "slice length is greater than or equal to min",
			minSize: MinSize{
				Min: 2,
			},
			input:    []int{1, 2, 3},
			expected: true,
		},
		{
			name: "slice length is less than min",
			minSize: MinSize{
				Min: 4,
			},
			input:    []int{1, 2, 3},
			expected: false,
		},
		{
			name: "input is not a string or slice",
			minSize: MinSize{
				Min: 5,
			},
			input:    123,
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

			result := tc.minSize.IsSatisfied(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
