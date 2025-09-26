package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_35(t *testing.T) {
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
	store.values[key] = value

	err := store.Delete(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if _, ok := store.values[key]; ok {
		t.Errorf("Expected key to be deleted, but it still exists")
	}
}
