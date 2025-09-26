package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIncr_1(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cache := NewMemoryCache()

	t.Run("Incr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "test_key"
		val := "10"
		err := cache.Put(ctx, key, val, time.Hour)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = cache.Incr(ctx, key)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		got, err := cache.Get(ctx, key)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		want := "11"
		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("Incr non-existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "non_existing_key"
		err := cache.Incr(ctx, key)
		if err != ErrKeyNotExist {
			t.Errorf("Expected error %v, got %v", ErrKeyNotExist, err)
		}
	})
}
