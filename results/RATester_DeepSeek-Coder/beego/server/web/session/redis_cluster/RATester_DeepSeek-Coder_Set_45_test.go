package redis_cluster

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_45(t *testing.T) {
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

	if store.values[key] != value {
		t.Errorf("Expected value to be %v, got %v", value, store.values[key])
	}
}
