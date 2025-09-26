package types

import (
	"fmt"
	"testing"
)

func TestAdd_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &EvictingStringQueue{
		size: 3,
		vals: []string{},
		set:  make(map[string]bool),
	}

	q.Add("a")
	q.Add("b")
	q.Add("c")

	if len(q.vals) != 3 {
		t.Errorf("Expected queue length to be 3, got %d", len(q.vals))
	}

	if !q.set["a"] || !q.set["b"] || !q.set["c"] {
		t.Errorf("Expected all elements to be in set")
	}

	q.Add("d")

	if len(q.vals) != 3 {
		t.Errorf("Expected queue length to be 3 after adding duplicate, got %d", len(q.vals))
	}

	if q.vals[0] != "b" || q.vals[1] != "c" || q.vals[2] != "d" {
		t.Errorf("Expected queue to be in LIFO order, got %v", q.vals)
	}

	if q.set["a"] {
		t.Errorf("Expected 'a' to be evicted")
	}
}
