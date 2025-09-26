package common

import (
	"fmt"
	"testing"
)

func TestInIntArr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []int{1, 2, 3, 4, 5}
	if !InIntArr(arr, 3) {
		t.Error("Expected true, got false")
	}
	if InIntArr(arr, 6) {
		t.Error("Expected false, got true")
	}
}
