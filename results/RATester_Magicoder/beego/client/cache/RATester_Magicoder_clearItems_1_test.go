package cache

import (
	"fmt"
	"testing"
)

func TestclearItems_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := MemoryCache{
		items: map[string]*MemoryItem{
			"key1": {val: "value1"},
			"key2": {val: "value2"},
		},
	}
	keys := []string{"key1", "key2"}
	cache.clearItems(keys)
	if len(cache.items) != 0 {
		t.Errorf("Expected cache items to be cleared, but found %d items", len(cache.items))
	}
}
