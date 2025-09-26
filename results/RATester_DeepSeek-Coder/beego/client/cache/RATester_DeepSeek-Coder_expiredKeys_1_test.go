package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestExpiredKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{
		items: map[string]*MemoryItem{
			"key1": {
				val:         "value1",
				createdTime: time.Now().Add(-2 * time.Hour),
				lifespan:    1 * time.Hour,
			},
			"key2": {
				val:         "value2",
				createdTime: time.Now().Add(-30 * time.Minute),
				lifespan:    1 * time.Hour,
			},
			"key3": {
				val:         "value3",
				createdTime: time.Now().Add(-50 * time.Minute),
				lifespan:    1 * time.Hour,
			},
		},
	}

	keys := bc.expiredKeys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 expired keys, got %d", len(keys))
	}

	for _, key := range keys {
		if key != "key1" && key != "key3" {
			t.Errorf("Unexpected key in expired keys: %s", key)
		}
	}
}
