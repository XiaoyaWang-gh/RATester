package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_36(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &CookieSessionStore{
		sid:    "test_sid",
		values: make(map[interface{}]interface{}),
	}

	key := "test_key"
	value := "test_value"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if store.values[key] != value {
		t.Errorf("Expected value to be set in store, but it was not")
	}
}
