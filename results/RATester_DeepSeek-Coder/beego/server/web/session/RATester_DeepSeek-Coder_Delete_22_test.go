package session

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	fs := &FileSessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}

	// Add some values to the session store
	fs.values["key1"] = "value1"
	fs.values["key2"] = "value2"

	// Delete a value
	err := fs.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check that the value was deleted
	if _, ok := fs.values["key1"]; ok {
		t.Errorf("Expected key1 to be deleted")
	}

	// Delete a non-existing value
	err = fs.Delete(ctx, "key3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check that the non-existing value was not deleted
	if _, ok := fs.values["key3"]; ok {
		t.Errorf("Expected key3 not to be deleted")
	}
}
