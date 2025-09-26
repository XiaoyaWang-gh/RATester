package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_38(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	fs := &FileSessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}

	key := "testKey"
	value := "testValue"

	err := fs.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fs.values[key] != value {
		t.Errorf("Expected value %v, but got %v", value, fs.values[key])
	}
}
