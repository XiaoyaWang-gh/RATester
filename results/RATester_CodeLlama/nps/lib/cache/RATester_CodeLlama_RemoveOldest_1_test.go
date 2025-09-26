package cache

import (
	"fmt"
	"testing"
)

func TestRemoveOldest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(2)
	c.Add("1", 1)
	c.Add("2", 2)
	c.Add("3", 3)
	if c.Len() != 2 {
		t.Errorf("cache length want 2 but got %d", c.Len())
	}
	c.RemoveOldest()
	if c.Len() != 1 {
		t.Errorf("cache length want 1 but got %d", c.Len())
	}
	c.RemoveOldest()
	if c.Len() != 0 {
		t.Errorf("cache length want 0 but got %d", c.Len())
	}
}
