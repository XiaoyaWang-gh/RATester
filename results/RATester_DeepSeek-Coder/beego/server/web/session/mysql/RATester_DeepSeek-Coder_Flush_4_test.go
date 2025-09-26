package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_4(t *testing.T) {
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
		t.Errorf("Expected no error, got %v", err)
	}

	if len(store.values) != 0 {
		t.Errorf("Expected values map to be empty, got %v", store.values)
	}
}
