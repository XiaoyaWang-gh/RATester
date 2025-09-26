package session

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &CookieSessionStore{
		values: make(map[interface{}]interface{}),
	}

	store.values["key1"] = "value1"
	store.values["key2"] = "value2"

	err := store.Delete(ctx, "key1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if _, ok := store.values["key1"]; ok {
		t.Errorf("Expected key1 to be deleted")
	}

	if _, ok := store.values["key2"]; !ok {
		t.Errorf("Expected key2 to still exist")
	}
}
