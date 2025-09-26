package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIncr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{
		items: map[string]*MemoryItem{
			"key1": {
				val:         "1",
				createdTime: time.Now(),
				lifespan:    10 * time.Second,
			},
		},
	}
	err := bc.Incr(context.Background(), "key1")
	if err != nil {
		t.Errorf("Incr() error = %v, want nil", err)
		return
	}
	if bc.items["key1"].val != "2" {
		t.Errorf("Incr() val = %v, want 2", bc.items["key1"].val)
	}
}
