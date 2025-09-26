package session

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &CookieSessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := store.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() error = %v, want nil", err)
		return
	}

	if len(store.values) != 0 {
		t.Errorf("Flush() len(store.values) = %v, want 0", len(store.values))
	}
}
