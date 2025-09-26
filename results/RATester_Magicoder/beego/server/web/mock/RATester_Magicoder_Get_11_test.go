package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}
	store.values["key"] = "value"

	result := store.Get(ctx, "key")
	if result != "value" {
		t.Errorf("Expected 'value', got '%v'", result)
	}
}
