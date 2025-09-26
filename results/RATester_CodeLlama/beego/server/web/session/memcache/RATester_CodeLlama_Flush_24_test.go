package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &SessionStore{}
	rs.Flush(context.Background())
	if len(rs.values) != 0 {
		t.Errorf("Flush() failed")
	}
}
