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

	// Add some values to the session
	fs.values["key1"] = "value1"
	fs.values["key2"] = "value2"

	// Delete a value
	err := fs.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the value was deleted
	if _, ok := fs.values["key1"]; ok {
		t.Error("Value was not deleted")
	}

	// Try to delete a non-existent value
	err = fs.Delete(ctx, "non-existent")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
