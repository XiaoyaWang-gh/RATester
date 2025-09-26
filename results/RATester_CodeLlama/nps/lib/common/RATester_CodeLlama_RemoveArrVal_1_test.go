package common

import (
	"fmt"
	"testing"
)

func TestRemoveArrVal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	arr := []string{"a", "b", "c", "d", "e"}
	arr = RemoveArrVal(arr, "b")
	if len(arr) != 4 {
		t.Errorf("RemoveArrVal failed")
	}
	if arr[0] != "a" || arr[1] != "c" || arr[2] != "d" || arr[3] != "e" {
		t.Errorf("RemoveArrVal failed")
	}
}
