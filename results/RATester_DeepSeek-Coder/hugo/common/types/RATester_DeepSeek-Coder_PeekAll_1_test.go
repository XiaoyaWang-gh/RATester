package types

import (
	"fmt"
	"testing"
)

func TestPeekAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &EvictingStringQueue{
		size: 5,
		vals: []string{"a", "b", "c", "d", "e"},
		set:  map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true},
	}

	peeked := q.PeekAll()

	if len(peeked) != 5 {
		t.Errorf("Expected length of peeked slice to be 5, got %d", len(peeked))
	}

	for i, v := range peeked {
		if v != q.vals[i] {
			t.Errorf("Expected value at index %d to be %s, got %s", i, q.vals[i], v)
		}
	}
}
