package common

import (
	"fmt"
	"testing"
)

func TestIsArrContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"a", "b", "c"}
	if !IsArrContains(arr, "a") {
		t.Error("IsArrContains failed")
	}
	if IsArrContains(arr, "d") {
		t.Error("IsArrContains failed")
	}
}
