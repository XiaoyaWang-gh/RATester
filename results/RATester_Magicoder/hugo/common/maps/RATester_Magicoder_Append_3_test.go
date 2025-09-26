package maps

import (
	"fmt"
	"testing"
)

func TestAppend_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &SliceCache[int]{m: make(map[string][]int)}
	cache.Append("key1", 1, 2, 3)
	cache.Append("key2", 4, 5, 6)

	value, ok := cache.Get("key1")
	if !ok || len(value) != 3 || value[0] != 1 || value[1] != 2 || value[2] != 3 {
		t.Errorf("Append failed for key1")
	}

	value, ok = cache.Get("key2")
	if !ok || len(value) != 3 || value[0] != 4 || value[1] != 5 || value[2] != 6 {
		t.Errorf("Append failed for key2")
	}
}
