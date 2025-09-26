package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_27(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		sid:    "test",
		values: make(map[interface{}]interface{}),
	}

	key := "key"
	value := "value"

	err := store.Set(ctx, key, value)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if store.values[key] != value {
		t.Errorf("Expected value to be %v, got %v", value, store.values[key])
	}
}
