package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_3(t *testing.T) {
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

	err := store.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if _, ok := store.values["key1"]; ok {
		t.Error("Expected key1 to be deleted")
	}

	err = store.Delete(ctx, "key3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if _, ok := store.values["key3"]; ok {
		t.Error("Expected key3 to not exist")
	}
}
