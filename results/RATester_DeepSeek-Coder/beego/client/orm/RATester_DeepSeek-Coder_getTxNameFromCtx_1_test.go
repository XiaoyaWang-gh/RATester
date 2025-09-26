package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestGetTxNameFromCtx_1(t *testing.T) {
	t.Run("should return empty string when context does not have TxNameKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		got := getTxNameFromCtx(ctx)
		want := ""
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return the value associated with TxNameKey when context has TxNameKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		want := "test"
		ctx := context.WithValue(context.Background(), TxNameKey, want)
		got := getTxNameFromCtx(ctx)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
