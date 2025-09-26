package hugo

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarkupScope_1(t *testing.T) {
	t.Run("should return markup scope from context", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		ctx = context.WithValue(ctx, "markupScope", "test")
		got := GetMarkupScope(ctx)
		want := "test"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return empty string if markup scope not in context", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		got := GetMarkupScope(ctx)
		want := ""
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
