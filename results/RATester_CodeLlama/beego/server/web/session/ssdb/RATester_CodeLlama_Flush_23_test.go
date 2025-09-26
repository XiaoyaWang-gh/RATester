package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{}
	s.Flush(context.Background())
	if s.values != nil {
		t.Errorf("Flush() failed")
	}
}
