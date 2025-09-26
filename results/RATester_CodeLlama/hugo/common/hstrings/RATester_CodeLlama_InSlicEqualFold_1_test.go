package hstrings

import (
	"fmt"
	"testing"
)

func TestInSlicEqualFold_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"a", "b", "c"}
	el := "A"
	if !InSlicEqualFold(arr, el) {
		t.Errorf("Expected %v to be in %v", el, arr)
	}
}
