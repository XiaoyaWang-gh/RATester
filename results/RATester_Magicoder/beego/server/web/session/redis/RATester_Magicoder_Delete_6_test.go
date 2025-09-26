package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	// Add some values to the store
	store.values["key1"] = "value1"
	store.values["key2"] = "value2"

	// Delete a value
	err := store.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the value was deleted
	_, ok := store.values["key1"]
	if ok {
		t.Error("Expected key1 to be deleted")
	}

	// Try to delete a non-existent value
	err = store.Delete(ctx, "key3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
