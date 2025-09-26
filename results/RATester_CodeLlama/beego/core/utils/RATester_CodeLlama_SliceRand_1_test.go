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

	a := []interface{}{1, 2, 3, 4, 5}
	b := SliceRand(a)
	if b != 1 && b != 2 && b != 3 && b != 4 && b != 5 {
		t.Errorf("SliceRand failed")
	}
}
