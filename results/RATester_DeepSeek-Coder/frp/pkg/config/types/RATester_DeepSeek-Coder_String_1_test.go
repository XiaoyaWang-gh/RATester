package types

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	type testCase struct {
		name     string
		input    PortsRangeSlice
		expected string
	}

	testCases := []testCase{
		{
			name: "Single port",
			input: PortsRangeSlice{
				{Single: 80},
			},
			expected: "80",
		},
		{
			name: "Range of ports",
			input: PortsRangeSlice{
				{Start: 80, End: 82},
			},
			expected: "80-82",
		},
		{
			name: "Multiple ports",
			input: PortsRangeSlice{
				{Single: 80},
				{Start: 8000, End: 8002},
				{Single: 443},
			},
			expected: "80,8000-8002,443",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
