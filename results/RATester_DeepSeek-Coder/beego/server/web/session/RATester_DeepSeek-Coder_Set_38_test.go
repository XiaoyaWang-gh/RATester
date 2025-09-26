package session

import (
	"context"
	"fmt"
	"sync"
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
		lock:   sync.RWMutex{},
		values: make(map[interface{}]interface{}),
	}

	key := "testKey"
	value := "testValue"

	err := fs.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if fs.values[key] != value {
		t.Errorf("Expected value to be %v, got %v", value, fs.values[key])
	}
}
