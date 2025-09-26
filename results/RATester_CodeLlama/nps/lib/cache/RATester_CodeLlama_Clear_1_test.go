package cache

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

	c := New(10)
	c.Add(1, 1)
	c.Add(2, 2)
	c.Add(3, 3)
	c.Clear()
	if c.Len() != 0 {
		t.Errorf("cache len = %d, want 0", c.Len())
	}
}
