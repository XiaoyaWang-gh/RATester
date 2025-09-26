package types

import (
	"fmt"
	"testing"
)

func TestLen_12(t *testing.T) {
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

	expected := 5
	actual := q.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
