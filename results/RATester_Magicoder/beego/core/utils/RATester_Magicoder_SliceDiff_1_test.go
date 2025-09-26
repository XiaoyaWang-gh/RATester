package utils

import (
	"fmt"
	"testing"
)

func TestSliceDiff_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []interface{}{1, 2, 3, 4, 5}
	slice2 := []interface{}{4, 5, 6, 7, 8}
	diffslice := SliceDiff(slice1, slice2)

	if len(diffslice) != 3 {
		t.Errorf("Expected length of diffslice to be 3, but got %d", len(diffslice))
	}

	for _, v := range diffslice {
		if v != 1 && v != 2 && v != 3 {
			t.Errorf("Unexpected value in diffslice: %v", v)
		}
	}
}
