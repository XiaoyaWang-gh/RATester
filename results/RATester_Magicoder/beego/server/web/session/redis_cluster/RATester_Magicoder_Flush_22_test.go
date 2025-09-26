package redis_cluster

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rs := &SessionStore{
		values: make(map[interface{}]interface{}),
	}

	// Add some values to the session
	rs.values["key1"] = "value1"
	rs.values["key2"] = "value2"

	// Flush the session
	err := rs.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}

	// Check if the session is empty
	if len(rs.values) != 0 {
		t.Errorf("Flush() did not clear the session values")
	}
}
