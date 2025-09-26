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

	fs := &FileSessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := fs.Flush(context.Background())
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}

	if len(fs.values) != 0 {
		t.Errorf("Flush() did not clear the values map. Expected: 0, Got: %d", len(fs.values))
	}
}
