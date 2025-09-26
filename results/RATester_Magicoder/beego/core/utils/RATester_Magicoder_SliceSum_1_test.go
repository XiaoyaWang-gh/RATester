package utils

import (
	"fmt"
	"testing"
)

func TestSliceSum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    []int64
		expected int64
	}

	tests := []test{
		{input: []int64{1, 2, 3, 4, 5}, expected: 15},
		{input: []int64{10, 20, 30, 40, 50}, expected: 150},
		{input: []int64{}, expected: 0},
		{input: []int64{-1, -2, -3, -4, -5}, expected: -15},
	}

	for _, test := range tests {
		result := SliceSum(test.input)
		if result != test.expected {
			t.Errorf("SliceSum(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
