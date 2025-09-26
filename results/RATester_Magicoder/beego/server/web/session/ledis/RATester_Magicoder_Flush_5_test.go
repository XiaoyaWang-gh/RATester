package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_5(t *testing.T) {
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

	// Call the method under test
	err := store.Flush(ctx)

	// Check that the error is nil
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Check that the values map is empty
	if len(store.values) != 0 {
		t.Errorf("Expected values map to be empty, got %v", store.values)
	}
}
