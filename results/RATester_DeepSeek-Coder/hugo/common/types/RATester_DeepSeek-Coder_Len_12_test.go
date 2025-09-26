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
		size: 3,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	got := q.Len()
	want := 3

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
