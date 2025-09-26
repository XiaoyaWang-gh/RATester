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
		t.Errorf("Flush() returned an error: %v", err)
	}

	if len(store.values) != 0 {
		t.Errorf("Flush() did not clear the session data, expected 0, got %d", len(store.values))
	}
}
