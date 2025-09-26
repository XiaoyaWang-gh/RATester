package maps

import (
	"fmt"
	"testing"
)

func TestMergeShallow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dst := map[string]any{
		"a": 1,
		"b": 2,
	}
	src := map[string]any{
		"c": 3,
		"d": 4,
	}
	MergeShallow(dst, src)
	if len(dst) != 4 {
		t.Errorf("dst length is %d, want 4", len(dst))
	}
	if dst["c"] != 3 {
		t.Errorf("dst[\"c\"] is %v, want 3", dst["c"])
	}
	if dst["d"] != 4 {
		t.Errorf("dst[\"d\"] is %v, want 4", dst["d"])
	}
}
