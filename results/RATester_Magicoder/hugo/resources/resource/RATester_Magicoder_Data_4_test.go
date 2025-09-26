package resource

import (
	"fmt"
	"testing"
)

func TestData_4(t *testing.T) {
	testCases := []struct {
		name     string
		input    *resourceError
		expected any
	}{
		{
			name: "Test case 1",
			input: &resourceError{
				data: "test data",
			},
			expected: "test data",
		},
		{
			name: "Test case 2",
			input: &resourceError{
				data: 123,
			},
			expected: 123,
		},
		{
			name: "Test case 3",
			input: &resourceError{
				data: nil,
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

			result := tc.input.Data()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
