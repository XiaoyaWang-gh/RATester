package types

import (
	"fmt"
	"testing"
)

func TestPeekAllSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &EvictingStringQueue{
		size: 10,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	set := q.PeekAllSet()

	if len(set) != 3 {
		t.Errorf("Expected 3 elements in set, got %d", len(set))
	}

	if !set["a"] {
		t.Errorf("Expected set to contain 'a'")
	}

	if !set["b"] {
		t.Errorf("Expected set to contain 'b'")
	}

	if !set["c"] {
		t.Errorf("Expected set to contain 'c'")
	}
}
