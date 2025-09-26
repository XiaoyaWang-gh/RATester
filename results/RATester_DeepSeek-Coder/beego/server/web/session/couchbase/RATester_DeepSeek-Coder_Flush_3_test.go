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
		values: make(map[interface{}]interface{}),
	}

	// Add some values to the session store
	cs.values["key1"] = "value1"
	cs.values["key2"] = "value2"

	// Flush the session store
	err := cs.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() error = %v, want nil", err)
		return
	}

	// Check that the values have been removed from the session store
	if len(cs.values) != 0 {
		t.Errorf("Flush() error = %v, want empty map", cs.values)
		return
	}
}
