package ping

import (
	"context"
	"fmt"
	"testing"
)

func TestWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Handler{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h.WithContext(ctx)
	if !h.terminating {
		t.Errorf("Handler.WithContext() failed to set terminating to true")
	}
}
