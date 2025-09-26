package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestCount_6(t *testing.T) {
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

	count := b.Count()
	if count != 3 {
		t.Errorf("Expected count to be 3, got %d", count)
	}
}
