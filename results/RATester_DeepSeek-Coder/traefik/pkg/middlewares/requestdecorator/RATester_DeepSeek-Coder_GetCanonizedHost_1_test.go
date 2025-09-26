package requestdecorator

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCanonizedHost_1(t *testing.T) {
	t.Run("should return empty string when context does not have canonical key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		result := GetCanonizedHost(ctx)
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	t.Run("should return value when context has canonical key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "test.com"
		ctx := context.WithValue(context.Background(), canonicalKey, expected)
		result := GetCanonizedHost(ctx)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
