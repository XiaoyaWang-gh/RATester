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

	t.Parallel()

	q := &EvictingStringQueue{
		size: 3,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	set := q.PeekAllSet()

	if len(set) != q.Len() {
		t.Errorf("expected len of set to be %d, got %d", q.Len(), len(set))
	}

	for _, v := range q.vals {
		if !set[v] {
			t.Errorf("expected set to contain %s", v)
		}
	}
}
