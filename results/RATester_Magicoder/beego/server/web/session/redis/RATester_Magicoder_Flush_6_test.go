package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rs := &SessionStore{
		values: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := rs.Flush(ctx)
	if err != nil {
		t.Errorf("Flush() returned an error: %v", err)
	}

	if len(rs.values) != 0 {
		t.Errorf("Flush() did not clear the values map. Expected: 0, Got: %d", len(rs.values))
	}
}
