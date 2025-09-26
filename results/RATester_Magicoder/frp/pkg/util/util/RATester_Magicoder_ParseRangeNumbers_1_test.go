package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseRangeNumbers_1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []int64
		err      error
	}

	testCases := []testCase{
		{
			name:     "single number",
			input:    "1",
			expected: []int64{1},
			err:      nil,
		},
		{
			name:     "range number",
			input:    "1-3",
			expected: []int64{1, 2, 3},
			err:      nil,
		},
		{
			name:     "invalid number",
			input:    "a-b",
			expected: nil,
			err:      fmt.Errorf("range number is invalid"),
		},
		{
			name:     "invalid range",
			input:    "3-1",
			expected: nil,
			err:      fmt.Errorf("range number is invalid"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			numbers, err := ParseRangeNumbers(tc.input)
			if !reflect.DeepEqual(numbers, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, numbers)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
