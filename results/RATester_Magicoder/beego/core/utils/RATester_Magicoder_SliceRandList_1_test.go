package utils

import (
	"fmt"
	"sort"
	"testing"
)

func TestSliceRandList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := 1
	max := 10
	list := SliceRandList(min, max)

	if len(list) != max-min+1 {
		t.Errorf("Expected length of list to be %d, but got %d", max-min+1, len(list))
	}

	for _, num := range list {
		if num < min || num > max {
			t.Errorf("Number %d is not within the range %d to %d", num, min, max)
		}
	}

	sort.Ints(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] == list[i] {
			t.Errorf("List contains duplicate numbers")
		}
	}
}
