package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestGet_26(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fs := &FileSessionStore{
		sid:    "test",
		lock:   sync.RWMutex{},
		values: make(map[interface{}]interface{}),
	}

	t.Run("get existing value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs.values["existingKey"] = "existingValue"
		value := fs.Get(ctx, "existingKey")
		if value != "existingValue" {
			t.Errorf("expected 'existingValue', got '%v'", value)
		}
	})

	t.Run("get non-existing value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := fs.Get(ctx, "nonExistingKey")
		if value != nil {
			t.Errorf("expected nil, got '%v'", value)
		}
	})
}
