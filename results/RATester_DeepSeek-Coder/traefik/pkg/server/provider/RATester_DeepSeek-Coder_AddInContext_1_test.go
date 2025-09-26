package provider

import (
	"context"
	"fmt"
	"testing"
)

func TestAddInContext_1(t *testing.T) {
	t.Run("should return the same context if the provider is the same", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		ctx = context.WithValue(ctx, key, "provider1")
		newCtx := AddInContext(ctx, "element@provider1")
		if newCtx != ctx {
			t.Errorf("Expected the same context, got %v and %v", newCtx, ctx)
		}
	})

	t.Run("should return a new context with the new provider", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		newCtx := AddInContext(ctx, "element@provider2")
		if newCtx == ctx {
			t.Errorf("Expected a new context, got the same context")
		}
		if newCtx.Value(key) != "provider2" {
			t.Errorf("Expected provider to be 'provider2', got %v", newCtx.Value(key))
		}
	})

	t.Run("should return the same context if the element name does not contain a provider", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		newCtx := AddInContext(ctx, "element")
		if newCtx != ctx {
			t.Errorf("Expected the same context, got %v and %v", newCtx, ctx)
		}
	})
}
