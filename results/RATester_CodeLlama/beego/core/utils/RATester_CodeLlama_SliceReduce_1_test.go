package utils

import (
	"fmt"
	"testing"
)

func TestSliceReduce_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice := []interface{}{1, 2, 3, 4, 5}
	a := func(v interface{}) interface{} {
		return v.(int) * 2
	}
	dslice := SliceReduce(slice, a)
	if len(dslice) != len(slice) {
		t.Errorf("SliceReduce() len(dslice) = %d, want %d", len(dslice), len(slice))
	}
	for i, v := range dslice {
		if v != a(slice[i]) {
			t.Errorf("SliceReduce() dslice[%d] = %v, want %v", i, v, a(slice[i]))
		}
	}
}
