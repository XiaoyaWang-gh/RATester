package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_9(t *testing.T) {
	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	t.Run("should return nil when key does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store.values["key1"] = "value1"
		store.values["key2"] = "value2"

		got := store.Get(ctx, "key3")
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("should return value when key exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store.values["key3"] = "value3"

		got := store.Get(ctx, "key3")
		if got != "value3" {
			t.Errorf("expected 'value3', got %v", got)
		}
	})
}
