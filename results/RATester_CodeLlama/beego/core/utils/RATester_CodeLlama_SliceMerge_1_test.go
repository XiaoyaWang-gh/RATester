package utils

import (
	"fmt"
	"testing"
)

func TestSliceMerge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []interface{}{1, 2, 3}
	c := SliceMerge(slice1, []interface{}{4, 5, 6})
	if len(c) != 6 {
		t.Errorf("SliceMerge failed")
	}
}
