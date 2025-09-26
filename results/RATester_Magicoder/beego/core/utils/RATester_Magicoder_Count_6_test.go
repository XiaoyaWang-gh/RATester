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

	bm := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	bm.Set("key1", "value1")
	bm.Set("key2", "value2")
	bm.Set("key3", "value3")

	count := bm.Count()
	if count != 3 {
		t.Errorf("Expected count to be 3, but got %d", count)
	}
}
