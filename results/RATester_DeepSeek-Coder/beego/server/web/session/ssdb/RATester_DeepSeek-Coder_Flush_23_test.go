package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_23(t *testing.T) {
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

	// Call the Flush method
	err := store.Flush(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check that the values have been flushed
	if len(store.values) != 0 {
		t.Errorf("Expected values to be flushed, got %v", store.values)
	}
}
