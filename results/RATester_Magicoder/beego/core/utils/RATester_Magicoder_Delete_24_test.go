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

	bm := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	bm.Set("key1", "value1")
	bm.Set("key2", "value2")

	bm.Delete("key1")

	if bm.Check("key1") {
		t.Error("Expected key1 to be deleted")
	}

	if !bm.Check("key2") {
		t.Error("Expected key2 to still exist")
	}
}
