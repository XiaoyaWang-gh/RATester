package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	// Add some values to the session store
	store.values["key1"] = "value1"
	store.values["key2"] = "value2"

	// Test Flush method
	err := store.Flush(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if all values are removed from the session store
	if len(store.values) != 0 {
		t.Errorf("Expected 0 values, got %d", len(store.values))
	}
}
