package common

import (
	"fmt"
	"testing"
)

func TestInStrArr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"a", "b", "c"}
	if !InStrArr(arr, "a") {
		t.Error("Expected true")
	}
	if InStrArr(arr, "d") {
		t.Error("Expected false")
	}
}
