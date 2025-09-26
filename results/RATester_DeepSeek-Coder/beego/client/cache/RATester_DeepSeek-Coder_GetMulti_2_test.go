package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMulti_2(t *testing.T) {
	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		keys := []string{"key1", "key2", "key3"}
		values := []interface{}{"value1", "value2", "value3"}

		for i, key := range keys {
			err := fc.Put(ctx, key, values[i], 0)
			if err != nil {
				t.Fatalf("failed to put key: %s, value: %v, error: %v", key, values[i], err)
			}
		}

		rc, err := fc.GetMulti(ctx, keys)
		if err != nil {
			t.Fatalf("failed to get multi keys: %v, error: %v", keys, err)
		}

		for i, value := range rc {
			if value != values[i] {
				t.Errorf("expected value: %v, got: %v", values[i], value)
			}
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		keys := []string{"key4", "key5", "key6"}
		_, err := fc.GetMulti(ctx, keys)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
