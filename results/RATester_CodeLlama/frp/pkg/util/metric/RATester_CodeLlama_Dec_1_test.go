package metric

import (
	"fmt"
	"testing"
)

func TestDec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &StandardDateCounter{
		reserveDays: 3,
		counts:      make([]int64, 3),
	}
	c.Inc(10)
	c.Dec(5)
	if c.counts[0] != 5 {
		t.Fatal("Dec failed")
	}
}
