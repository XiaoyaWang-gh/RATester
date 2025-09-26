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
		size: 3,
		vals: []string{"a", "b", "c", "d", "e"},
		set:  map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true},
	}
	vals := q.PeekAll()
	if len(vals) != 3 {
		t.Errorf("Expected len(vals) == 3, got %d", len(vals))
	}
	if vals[0] != "c" {
		t.Errorf("Expected vals[0] == 'c', got %q", vals[0])
	}
	if vals[1] != "d" {
		t.Errorf("Expected vals[1] == 'd', got %q", vals[1])
	}
	if vals[2] != "e" {
		t.Errorf("Expected vals[2] == 'e', got %q", vals[2])
	}
}
