package migration

import (
	"fmt"
	"testing"
)

func TestSwap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := dataSlice{data{1, "test1", nil}, data{2, "test2", nil}}
	d.Swap(0, 1)
	if d[0].created != 2 || d[1].created != 1 {
		t.Errorf("Swap failed. Expected: %v, %v. Got: %v, %v", 2, 1, d[0].created, d[1].created)
	}
}
