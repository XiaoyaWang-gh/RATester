package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_44(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		values: make(map[interface{}]interface{}),
	}
	key := "testKey"
	value := "testValue"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	val, ok := store.values[key]
	if !ok {
		t.Errorf("Expected key to be in values map")
	}

	if val != value {
		t.Errorf("Expected value to be %v, got %v", value, val)
	}
}
