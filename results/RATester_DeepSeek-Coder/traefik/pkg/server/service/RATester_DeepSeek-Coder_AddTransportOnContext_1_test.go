package service

import (
	"context"
	"fmt"
	"testing"
)

func TestAddTransportOnContext_1(t *testing.T) {
	t.Run("should add transport to context", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		newCtx := AddTransportOnContext(ctx)

		val := newCtx.Value(transportKey)
		if val == nil {
			t.Errorf("expected transport to be added to context, but it was not")
		}

		transport, ok := val.(*stickyRoundTripper)
		if !ok {
			t.Errorf("expected transport to be of type *stickyRoundTripper, but it was not")
		}

		if transport.RoundTripper != nil {
			t.Errorf("expected RoundTripper to be nil, but it was not")
		}
	})
}
