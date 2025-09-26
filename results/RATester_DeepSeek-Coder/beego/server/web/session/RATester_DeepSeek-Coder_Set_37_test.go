package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_37(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &MemSessionStore{
		value: make(map[interface{}]interface{}),
	}
	key := "testKey"
	value := "testValue"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if store.value[key] != value {
		t.Errorf("Expected value to be %v, got %v", value, store.value[key])
	}
}
