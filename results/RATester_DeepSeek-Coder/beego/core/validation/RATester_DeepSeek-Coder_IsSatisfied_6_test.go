package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_6(t *testing.T) {
	type testCase struct {
		name     string
		min      Min
		obj      interface{}
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test int64 value",
			min: Min{
				Min: 10,
			},
			obj:      int64(15),
			expected: true,
		},
		{
			name: "Test int value",
			min: Min{
				Min: 10,
			},
			obj:      int(5),
			expected: false,
		},
		{
			name: "Test int32 value",
			min: Min{
				Min: 10,
			},
			obj:      int32(15),
			expected: true,
		},
		{
			name: "Test int16 value",
			min: Min{
				Min: 10,
			},
			obj:      int16(5),
			expected: false,
		},
		{
			name: "Test int8 value",
			min: Min{
				Min: 10,
			},
			obj:      int8(5),
			expected: false,
		},
		{
			name: "Test unsupported type",
			min: Min{
				Min: 10,
			},
			obj:      "unsupported",
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

			result := tc.min.IsSatisfied(tc.obj)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
