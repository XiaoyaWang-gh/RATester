package plugin

import (
	"context"
	"fmt"
	"testing"
)

func TestGetReqidFromContext_1(t *testing.T) {
	t.Run("should return empty string when context does not have reqid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		got := GetReqidFromContext(ctx)
		want := ""
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return reqid when context has reqid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.WithValue(context.Background(), reqidKey, "test-reqid")
		got := GetReqidFromContext(ctx)
		want := "test-reqid"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
