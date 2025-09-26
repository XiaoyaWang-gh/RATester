package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := store.Flush(ctx)
	if err != nil {
		t.Errorf("Flush returned an error: %v", err)
	}

	if len(store.values) != 0 {
		t.Errorf("Flush did not clear the session store, expected 0 items, got %d", len(store.values))
	}
}
