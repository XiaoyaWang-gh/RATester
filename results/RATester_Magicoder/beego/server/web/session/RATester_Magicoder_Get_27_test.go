package session

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_27(t *testing.T) {
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

	store.values["test_key"] = "test_value"

	value := store.Get(ctx, "test_key")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%v'", value)
	}

	value = store.Get(ctx, "non_existent_key")
	if value != nil {
		t.Errorf("Expected nil, got '%v'", value)
	}
}
