package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestDecr_3(t *testing.T) {
	ctx := context.Background()
	cache := NewMemoryCache()

	t.Run("decrement existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "test_key"
		val := "10"
		cache.Put(ctx, key, val, 0)

		err := cache.Decr(ctx, key)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		got, err := cache.Get(ctx, key)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		want := "9"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("decrement non-existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "non_existing_key"
		err := cache.Decr(ctx, key)
		if err != ErrKeyNotExist {
			t.Errorf("got %v, want %v", err, ErrKeyNotExist)
		}
	})

	t.Run("decrement non-integer value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "non_integer_key"
		val := "not_an_integer"
		cache.Put(ctx, key, val, 0)

		err := cache.Decr(ctx, key)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
