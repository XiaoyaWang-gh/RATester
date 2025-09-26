package session

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	fs := &FileSessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := fs.Flush(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(fs.values) != 0 {
		t.Errorf("Expected values map to be empty, got %v", fs.values)
	}
}
