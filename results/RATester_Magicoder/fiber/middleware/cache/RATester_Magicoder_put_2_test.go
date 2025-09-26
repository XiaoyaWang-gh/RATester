package cache

import (
	"fmt"
	"testing"
)

func Testput_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{
		entries: []heapEntry{},
		indices: []int{},
		maxidx:  0,
	}

	key := "testKey"
	exp := uint64(100)
	bytes := uint(10)

	idx := h.put(key, exp, bytes)

	if idx != 0 {
		t.Errorf("Expected idx to be 0, but got %d", idx)
	}

	if len(h.entries) != 1 {
		t.Errorf("Expected entries length to be 1, but got %d", len(h.entries))
	}

	if h.entries[0].key != key {
		t.Errorf("Expected key to be %s, but got %s", key, h.entries[0].key)
	}

	if h.entries[0].exp != exp {
		t.Errorf("Expected exp to be %d, but got %d", exp, h.entries[0].exp)
	}

	if h.entries[0].bytes != bytes {
		t.Errorf("Expected bytes to be %d, but got %d", bytes, h.entries[0].bytes)
	}

	if h.entries[0].idx != 0 {
		t.Errorf("Expected idx to be 0, but got %d", h.entries[0].idx)
	}

	if len(h.indices) != 1 {
		t.Errorf("Expected indices length to be 1, but got %d", len(h.indices))
	}

	if h.indices[0] != 0 {
		t.Errorf("Expected index in indices to be 0, but got %d", h.indices[0])
	}

	if h.maxidx != 1 {
		t.Errorf("Expected maxidx to be 1, but got %d", h.maxidx)
	}
}
