package service

import (
	"context"
	"fmt"
	"testing"
)

func TestAddTransportOnContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	newCtx := AddTransportOnContext(ctx)

	val := newCtx.Value(transportKey)
	if val == nil {
		t.Errorf("Expected value to be set in context")
	}

	transport, ok := val.(*stickyRoundTripper)
	if !ok {
		t.Errorf("Expected value to be of type *stickyRoundTripper")
	}

	if transport.RoundTripper == nil {
		t.Errorf("Expected RoundTripper to be set in stickyRoundTripper")
	}
}
