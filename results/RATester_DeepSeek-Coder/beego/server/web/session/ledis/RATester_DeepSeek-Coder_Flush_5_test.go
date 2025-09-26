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

	// Test Flush
	err := store.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}

	if len(store.values) != 0 {
		t.Errorf("Flush() did not clear the values map. It still contains %d elements", len(store.values))
	}
}
