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

	intslice := []int64{1, 2, 3, 4, 5}
	sum := SliceSum(intslice)
	if sum != 15 {
		t.Errorf("SliceSum() = %v, want %v", sum, 15)
	}
}
