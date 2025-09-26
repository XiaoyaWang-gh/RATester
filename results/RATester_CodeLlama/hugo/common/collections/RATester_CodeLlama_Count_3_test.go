package collections

import (
	"fmt"
	"testing"
)

func TestCount_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ss := SortedStringSlice{"a", "b", "c", "d", "e"}
	if ss.Count("a") != 1 {
		t.Errorf("Count(\"a\") = %d, want 1", ss.Count("a"))
	}
	if ss.Count("b") != 1 {
		t.Errorf("Count(\"b\") = %d, want 1", ss.Count("b"))
	}
	if ss.Count("c") != 1 {
		t.Errorf("Count(\"c\") = %d, want 1", ss.Count("c"))
	}
	if ss.Count("d") != 1 {
		t.Errorf("Count(\"d\") = %d, want 1", ss.Count("d"))
	}
	if ss.Count("e") != 1 {
		t.Errorf("Count(\"e\") = %d, want 1", ss.Count("e"))
	}
	if ss.Count("f") != 0 {
		t.Errorf("Count(\"f\") = %d, want 0", ss.Count("f"))
	}
}
