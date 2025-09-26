package maps

import (
	"fmt"
	"testing"
)

func TestDrain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache[string, string]{
		m: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	}

	m := c.Drain()

	if len(m) != 3 {
		t.Errorf("expected 3, got %d", len(m))
	}

	if m["a"] != "1" {
		t.Errorf("expected 1, got %s", m["a"])
	}

	if m["b"] != "2" {
		t.Errorf("expected 2, got %s", m["b"])
	}

	if m["c"] != "3" {
		t.Errorf("expected 3, got %s", m["c"])
	}
}
