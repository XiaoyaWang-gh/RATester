package maps

import (
	"fmt"
	"testing"
)

func TestLookupEqualFold_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	key := "A"
	v, k, found := LookupEqualFold(m, key)
	if !found {
		t.Errorf("LookupEqualFold(%v, %v) = %v, %v, %v; want _, _, true", m, key, v, k, found)
	}
	if v != 1 {
		t.Errorf("LookupEqualFold(%v, %v) = %v, %v, %v; want 1, _, true", m, key, v, k, found)
	}
	if k != "a" {
		t.Errorf("LookupEqualFold(%v, %v) = %v, %v, %v; want _, a, true", m, key, v, k, found)
	}
}
