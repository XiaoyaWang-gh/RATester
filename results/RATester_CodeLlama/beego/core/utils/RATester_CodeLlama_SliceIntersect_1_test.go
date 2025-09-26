package utils

import (
	"fmt"
	"testing"
)

func TestSliceIntersect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []interface{}{1, 2, 3, 4, 5}
	diffslice := SliceIntersect(slice1, []interface{}{1, 2, 3})
	if len(diffslice) != 3 {
		t.Errorf("SliceIntersect() = %v, want %v", diffslice, 3)
	}
}
