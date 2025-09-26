package cache

import (
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &indexedHeap{}
	h.Push(heapEntry{key: "test", exp: 1, bytes: 1, idx: 1})
	if h.Len() != 1 {
		t.Errorf("Expected heap length to be 1, got %d", h.Len())
	}
	if h.entries[0].key != "test" {
		t.Errorf("Expected first entry key to be 'test', got %s", h.entries[0].key)
	}
	if h.entries[0].exp != 1 {
		t.Errorf("Expected first entry exp to be 1, got %d", h.entries[0].exp)
	}
	if h.entries[0].bytes != 1 {
		t.Errorf("Expected first entry bytes to be 1, got %d", h.entries[0].bytes)
	}
	if h.entries[0].idx != 1 {
		t.Errorf("Expected first entry idx to be 1, got %d", h.entries[0].idx)
	}
}
