package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelete_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	b.Set("key1", "value1")
	b.Set("key2", "value2")
	b.Set("key3", "value3")

	b.Delete("key2")

	if b.Count() != 2 {
		t.Errorf("Expected count after delete is 2, got %d", b.Count())
	}

	if b.Get("key2") != nil {
		t.Errorf("Expected nil after delete, got %v", b.Get("key2"))
	}
}
