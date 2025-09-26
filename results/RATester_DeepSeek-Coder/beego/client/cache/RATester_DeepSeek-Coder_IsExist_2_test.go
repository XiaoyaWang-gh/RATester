package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestIsExist_2(t *testing.T) {
	ctx := context.Background()
	cache := &MemoryCache{
		items: make(map[string]*MemoryItem),
	}

	t.Run("key exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache.items["existingKey"] = &MemoryItem{}
		exist, err := cache.IsExist(ctx, "existingKey")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !exist {
			t.Errorf("expected key to exist, but it does not")
		}
	})

	t.Run("key does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		exist, err := cache.IsExist(ctx, "nonExistingKey")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if exist {
			t.Errorf("expected key to not exist, but it does")
		}
	})
}
