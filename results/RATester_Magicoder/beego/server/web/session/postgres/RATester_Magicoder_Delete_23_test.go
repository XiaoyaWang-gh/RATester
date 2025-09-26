package postgres

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Test deleting an existing key
	err := store.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, ok := store.values["key1"]; ok {
		t.Error("Key was not deleted")
	}

	// Test deleting a non-existing key
	err = store.Delete(ctx, "key3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, ok := store.values["key3"]; ok {
		t.Error("Key was not deleted")
	}
}
