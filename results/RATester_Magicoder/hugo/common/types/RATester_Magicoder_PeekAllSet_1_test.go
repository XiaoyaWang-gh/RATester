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

	q := NewEvictingStringQueue(3)
	q.Add("a")
	q.Add("b")
	q.Add("c")
	q.Add("d")

	set := q.PeekAllSet()

	if len(set) != 3 {
		t.Errorf("Expected set length to be 3, got %d", len(set))
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

	if set["d"] {
		t.Errorf("Expected set not to contain 'd'")
	}
}
