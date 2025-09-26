package session

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &MemSessionStore{
		value: make(map[interface{}]interface{}),
	}

	store.value["key1"] = "value1"
	store.value["key2"] = "value2"

	err := store.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, ok := store.value["key1"]; ok {
		t.Errorf("Expected key1 to be deleted")
	}

	if _, ok := store.value["key2"]; !ok {
		t.Errorf("Expected key2 to still exist")
	}
}
