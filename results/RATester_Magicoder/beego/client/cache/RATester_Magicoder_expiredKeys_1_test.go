package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestexpiredKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &MemoryCache{
		dur:   1 * time.Minute,
		items: make(map[string]*MemoryItem),
	}

	// Add some items with different expiration times
	cache.items["expired"] = &MemoryItem{
		val:         "expired",
		createdTime: time.Now().Add(-2 * time.Minute),
		lifespan:    1 * time.Minute,
	}
	cache.items["not_expired"] = &MemoryItem{
		val:         "not_expired",
		createdTime: time.Now().Add(-10 * time.Second),
		lifespan:    1 * time.Minute,
	}

	// Test expiredKeys function
	expiredKeys := cache.expiredKeys()
	if len(expiredKeys) != 1 || expiredKeys[0] != "expired" {
		t.Errorf("Expected expiredKeys to return ['expired'], but got %v", expiredKeys)
	}
}
