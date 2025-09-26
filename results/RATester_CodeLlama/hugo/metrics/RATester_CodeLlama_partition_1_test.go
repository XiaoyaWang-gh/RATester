package metrics

import (
	"fmt"
	"testing"
)

func TestPartition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := 10
	scale := 2
	want := 8
	got := partition(d, scale)
	if got != want {
		t.Errorf("partition(%d, %d) = %d; want %d", d, scale, got, want)
	}
}
