package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_37(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &MemSessionStore{
		value: make(map[interface{}]interface{}),
	}

	ctx := context.Background()
	key := "testKey"
	value := "testValue"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	storedValue, ok := store.value[key]
	if !ok {
		t.Errorf("Expected key %v to be in store", key)
	}

	if storedValue != value {
		t.Errorf("Expected value %v, but got %v", value, storedValue)
	}
}
