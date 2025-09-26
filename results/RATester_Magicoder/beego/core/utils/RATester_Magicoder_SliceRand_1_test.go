package utils

import (
	"fmt"
	"testing"
)

func TestSliceRand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := []interface{}{1, "two", 3.0, "four", 5}
	b := SliceRand(a)

	// Check if b is in a
	found := false
	for _, v := range a {
		if v == b {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("SliceRand() returned value not in original slice")
	}
}
