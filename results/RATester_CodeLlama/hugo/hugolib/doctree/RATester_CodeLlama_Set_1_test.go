package doctree

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := DimensionFlag(0)
	o := DimensionFlag(1)
	if got := d.Set(o); got != DimensionFlag(1) {
		t.Errorf("DimensionFlag.Set() = %v, want %v", got, DimensionFlag(1))
	}
}
