package cache

import (
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(2)
	c.Add("a", 1)
	c.Add("b", 2)
	c.Add("c", 3)
	c.Remove("b")
	if v, ok := c.Get("b"); ok {
		t.Errorf("Get(b) = %v, %v, want %v, %v", v, ok, nil, false)
	}
	if c.Len() != 2 {
		t.Errorf("cache.Len() = %v, want %v", c.Len(), 2)
	}
	c.Remove("a")
	c.Remove("c")
	if c.Len() != 0 {
		t.Errorf("cache.Len() = %v, want %v", c.Len(), 0)
	}
}
