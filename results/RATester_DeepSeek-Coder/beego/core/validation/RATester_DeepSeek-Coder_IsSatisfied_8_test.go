package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_8(t *testing.T) {
	type testCase struct {
		name     string
		max      Max
		input    interface{}
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Test with int64 input",
			max:      Max{Max: 10},
			input:    int64(5),
			expected: true,
		},
		{
			name:     "Test with int input",
			max:      Max{Max: 10},
			input:    5,
			expected: true,
		},
		{
			name:     "Test with int32 input",
			max:      Max{Max: 10},
			input:    int32(5),
			expected: true,
		},
		{
			name:     "Test with int16 input",
			max:      Max{Max: 10},
			input:    int16(5),
			expected: true,
		},
		{
			name:     "Test with int8 input",
			max:      Max{Max: 10},
			input:    int8(5),
			expected: true,
		},
		{
			name:     "Test with float64 input",
			max:      Max{Max: 10},
			input:    float64(5),
			expected: false,
		},
		{
			name:     "Test with string input",
			max:      Max{Max: 10},
			input:    "5",
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

			result := tc.max.IsSatisfied(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
