package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cs := &SessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := cs.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}

	if len(cs.values) != 0 {
		t.Errorf("Flush() did not clear the session store")
	}
}
