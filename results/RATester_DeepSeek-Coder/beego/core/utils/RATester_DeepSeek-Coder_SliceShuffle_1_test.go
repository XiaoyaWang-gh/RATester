package utils

import (
	"fmt"
	"testing"
)

func TestSliceShuffle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := []interface{}{1, 2, 3, 4, 5}
	shuffled := SliceShuffle(slice)

	if len(shuffled) != len(slice) {
		t.Errorf("Expected length of shuffled slice to be %d, got %d", len(slice), len(shuffled))
	}

	same := true
	for i := 0; i < len(slice); i++ {
		if slice[i] != shuffled[i] {
			same = false
			break
		}
	}

	if same {
		t.Errorf("Expected shuffled slice to be different from original slice")
	}
}
