package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_42(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	// Test case 1: key exists
	store.values["key1"] = "value1"
	value := store.Get(ctx, "key1")
	if value != "value1" {
		t.Errorf("Expected 'value1', got '%v'", value)
	}

	// Test case 2: key does not exist
	value = store.Get(ctx, "key2")
	if value != nil {
		t.Errorf("Expected nil, got '%v'", value)
	}
}
