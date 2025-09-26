package capture

import (
	"context"
	"fmt"
	"testing"
)

func TestFromContext_1(t *testing.T) {
	t.Run("should return error if value not found in context", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		_, err := FromContext(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error if value stored in context is not a *Capture", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.WithValue(context.Background(), capturedData, "not a Capture")
		_, err := FromContext(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return Capture if value stored in context is a *Capture", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		capt := &Capture{}
		ctx := context.WithValue(context.Background(), capturedData, capt)
		got, err := FromContext(ctx)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != *capt {
			t.Errorf("expected %v, got %v", *capt, got)
		}
	})
}
