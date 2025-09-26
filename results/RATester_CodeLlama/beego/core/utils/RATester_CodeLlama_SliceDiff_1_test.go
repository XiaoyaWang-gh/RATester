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
	diffslice := SliceDiff(slice1, []interface{}{1, 2, 3})
	if len(diffslice) != 2 {
		t.Errorf("SliceDiff() = %v, want %v", diffslice, []interface{}{4, 5})
	}
}
