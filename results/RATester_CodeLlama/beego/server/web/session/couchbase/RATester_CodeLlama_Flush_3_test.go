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

	cs := &SessionStore{}
	cs.Flush(context.Background())
}
