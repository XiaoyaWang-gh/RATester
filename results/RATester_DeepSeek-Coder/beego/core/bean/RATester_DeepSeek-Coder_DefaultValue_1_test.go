package bean

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDefaultValue_1(t *testing.T) {
	ctx := context.Background()
	adapter := &TimeTypeAdapter{Layout: "2006-01-02 15:04:05"}

	t.Run("now", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val, err := adapter.DefaultValue(ctx, "now")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if val.(time.Time).IsZero() {
			t.Errorf("expected non-zero time, got zero time")
		}
	})

	t.Run("specific time", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val, err := adapter.DefaultValue(ctx, "2022-01-02 15:04:05")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected, _ := time.Parse("2006-01-02 15:04:05", "2022-01-02 15:04:05")
		if !val.(time.Time).Equal(expected) {
			t.Errorf("expected %v, got %v", expected, val)
		}
	})

	t.Run("invalid time", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := adapter.DefaultValue(ctx, "invalid")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
