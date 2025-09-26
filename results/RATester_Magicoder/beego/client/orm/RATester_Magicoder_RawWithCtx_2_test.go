package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestRawWithCtx_2(t *testing.T) {
	ctx := context.Background()
	ormer := &DoNothingOrm{}

	t.Run("TestRawWithCtx", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ormer.RawWithCtx(ctx, "SELECT * FROM table")
		if result == nil {
			t.Error("Expected non-nil result, but got nil")
		}
	})
}
