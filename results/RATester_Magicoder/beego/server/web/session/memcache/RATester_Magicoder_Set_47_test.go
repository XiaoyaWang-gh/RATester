package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_47(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	key := "testKey"
	value := "testValue"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	storedValue, ok := store.values[key]
	if !ok {
		t.Errorf("Expected key %v to be in the store", key)
	}

	if storedValue != value {
		t.Errorf("Expected value %v, but got %v", value, storedValue)
	}
}
