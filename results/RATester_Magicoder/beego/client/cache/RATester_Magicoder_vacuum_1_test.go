package cache

import (
	"fmt"
	"testing"
	"time"
)

func Testvacuum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{
		dur:   time.Second,
		Every: 1,
		items: map[string]*MemoryItem{
			"key1": {
				val:         "value1",
				createdTime: time.Now(),
				lifespan:    time.Hour,
			},
			"key2": {
				val:         "value2",
				createdTime: time.Now().Add(-time.Hour),
				lifespan:    time.Hour,
			},
		},
	}

	bc.vacuum()

	if len(bc.items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(bc.items))
	}
}
