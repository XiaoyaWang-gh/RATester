package pageparser

import (
	"fmt"
	"testing"
)

func TestMinIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{0, -1, 2, 3, 4}, 0},
		{[]int{-1, -2, -3, -4, -5}, -1},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{-1}, -1},
		{[]int{}, -1},
	}

	for _, tc := range tests {
		got := minIndex(tc.input...)
		if got != tc.want {
			t.Errorf("minIndex(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
