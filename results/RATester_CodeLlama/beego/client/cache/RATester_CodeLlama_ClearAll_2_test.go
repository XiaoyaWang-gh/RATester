package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestClearAll_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bc := &MemoryCache{}
	bc.items = map[string]*MemoryItem{
		"key1": {
			val:         "value1",
			createdTime: time.Now(),
			lifespan:    10 * time.Second,
		},
		"key2": {
			val:         "value2",
			createdTime: time.Now(),
			lifespan:    10 * time.Second,
		},
	}
	err := bc.ClearAll(context.Background())
	if err != nil {
		t.Errorf("ClearAll() error = %v, want nil", err)
	}
	if len(bc.items) != 0 {
		t.Errorf("ClearAll() items = %v, want 0", len(bc.items))
	}
}
