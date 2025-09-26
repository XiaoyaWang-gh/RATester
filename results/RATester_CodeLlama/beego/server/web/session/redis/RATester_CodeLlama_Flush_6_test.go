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

	rs := &SessionStore{}
	rs.Flush(context.Background())
}
