package metric

import (
	"fmt"
	"testing"
)

func TestClear_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &StandardDateCounter{
		reserveDays: 3,
		counts:      []int64{1, 2, 3},
	}
	c.Clear()
	if c.counts[0] != 0 || c.counts[1] != 0 || c.counts[2] != 0 {
		t.Fatal("clear failed")
	}
}
