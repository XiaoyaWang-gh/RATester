package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceShuffle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := []interface{}{1, 2, 3, 4, 5}
	shuffledSlice := SliceShuffle(slice)

	// Check if the shuffled slice is different from the original one
	if reflect.DeepEqual(slice, shuffledSlice) {
		t.Error("SliceShuffle function is not shuffling the slice")
	}

	// Check if the shuffled slice contains the same elements as the original one
	for _, v := range slice {
		found := false
		for _, v2 := range shuffledSlice {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			t.Error("SliceShuffle function is not shuffling the slice")
		}
	}
}
