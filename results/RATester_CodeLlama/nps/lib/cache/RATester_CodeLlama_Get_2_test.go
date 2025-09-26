package cache

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(2)
	c.Add(1, 1)
	c.Add(2, 2)
	c.Add(3, 3)
	if v, ok := c.Get(1); !ok || v != 1 {
		t.Error("cache[1] =", v, "expected 1")
	}
	if v, ok := c.Get(2); !ok || v != 2 {
		t.Error("cache[2] =", v, "expected 2")
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Error("cache[3] =", v, "expected 3")
	}
	if v, ok := c.Get(4); ok {
		t.Error("cache[4] =", v, "expected none")
	}
}
