package cache

import (
	"fmt"
	"testing"
)

func TestPushInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &indexedHeap{}
	entry := heapEntry{key: "key", exp: 1, bytes: 1, idx: 1}
	// Act
	h.pushInternal(entry)
	// Assert
	if len(h.entries) != 1 {
		t.Errorf("Expected len(h.entries) == 1, got %d", len(h.entries))
	}
	if h.entries[0].key != "key" {
		t.Errorf("Expected h.entries[0].key == 'key', got %s", h.entries[0].key)
	}
	if h.entries[0].exp != 1 {
		t.Errorf("Expected h.entries[0].exp == 1, got %d", h.entries[0].exp)
	}
	if h.entries[0].bytes != 1 {
		t.Errorf("Expected h.entries[0].bytes == 1, got %d", h.entries[0].bytes)
	}
	if h.entries[0].idx != 1 {
		t.Errorf("Expected h.entries[0].idx == 1, got %d", h.entries[0].idx)
	}
	if len(h.indices) != 1 {
		t.Errorf("Expected len(h.indices) == 1, got %d", len(h.indices))
	}
	if h.indices[1] != 0 {
		t.Errorf("Expected h.indices[1] == 0, got %d", h.indices[1])
	}
}
