package maps

import (
	"fmt"
	"testing"
)

func TestMergeParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dst := Params{"a": 1, "b": 2}
	src := Params{"a": 3, "c": 4}
	MergeParams(dst, src)
	if dst["a"] != 3 {
		t.Errorf("dst[\"a\"] = %v, want 3", dst["a"])
	}
	if dst["b"] != 2 {
		t.Errorf("dst[\"b\"] = %v, want 2", dst["b"])
	}
	if dst["c"] != 4 {
		t.Errorf("dst[\"c\"] = %v, want 4", dst["c"])
	}
}
