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

	rs := &SessionStore{}
	rs.values = make(map[interface{}]interface{})
	rs.values["key"] = "value"
	if err := rs.Flush(context.Background()); err != nil {
		t.Errorf("Flush() error = %v, want nil", err)
	}
	if len(rs.values) != 0 {
		t.Errorf("Flush() values = %v, want 0", len(rs.values))
	}
}
