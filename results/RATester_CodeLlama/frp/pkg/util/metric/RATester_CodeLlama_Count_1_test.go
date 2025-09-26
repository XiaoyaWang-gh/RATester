package metric

import (
	"fmt"
	"testing"
)

func TestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &StandardCounter{}
	c.Inc(1)
	if c.Count() != 1 {
		t.Error("Count() = ", c.Count())
	}
}
