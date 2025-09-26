package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_32(t *testing.T) {
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
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the value was deleted
	if _, ok := store.values["key1"]; ok {
		t.Errorf("Expected key1 to be deleted")
	}

	// Try to delete a non-existent value
	err = store.Delete(ctx, "key3")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the non-existent value was not deleted
	if _, ok := store.values["key3"]; ok {
		t.Errorf("Expected key3 not to be deleted")
	}
}
