package cache

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(2)
	c.Add(1, 1)
	c.Add(2, 2)
	c.Add(3, 3)
	if c.Len() != 2 {
		t.Errorf("cache length = %d, want 2", c.Len())
	}
	if _, ok := c.Get(1); ok {
		t.Errorf("cache should not contain 1")
	}
	if v, ok := c.Get(2); !ok || v != 2 {
		t.Errorf("cache should contain 2 = %v, want 2", v)
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Errorf("cache should contain 3 = %v, want 3", v)
	}
}
