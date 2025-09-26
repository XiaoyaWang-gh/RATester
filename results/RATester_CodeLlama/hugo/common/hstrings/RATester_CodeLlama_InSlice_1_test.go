package hstrings

import (
	"fmt"
	"testing"
)

func TestInSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"a", "b", "c"}
	el := "a"
	if !InSlice(arr, el) {
		t.Errorf("Expected %v to be in %v", el, arr)
	}
}
