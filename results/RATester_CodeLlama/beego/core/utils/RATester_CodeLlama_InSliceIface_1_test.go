package utils

import (
	"fmt"
	"testing"
)

func TestInSliceIface_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := 1
	sl := []interface{}{1, 2, 3}
	if !InSliceIface(v, sl) {
		t.Errorf("InSliceIface failed")
	}
}
