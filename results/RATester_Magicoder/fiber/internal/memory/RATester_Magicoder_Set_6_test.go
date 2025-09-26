package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storage := &Storage{
		data: make(map[string]item),
	}

	key := "testKey"
	val := "testValue"
	ttl := time.Second * 10

	storage.Set(key, val, ttl)

	storage.RLock()
	defer storage.RUnlock()

	item, ok := storage.data[key]
	if !ok {
		t.Errorf("Expected key %s to exist in storage", key)
	}

	if item.v != val {
		t.Errorf("Expected value %v, got %v", val, item.v)
	}

	if item.e < uint32(time.Now().Unix()) {
		t.Errorf("Expected expiration time to be in the future, got %v", item.e)
	}
}
