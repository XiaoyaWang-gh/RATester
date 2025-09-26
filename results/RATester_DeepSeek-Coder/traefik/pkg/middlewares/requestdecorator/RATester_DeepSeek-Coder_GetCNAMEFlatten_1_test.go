package requestdecorator

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCNAMEFlatten_1(t *testing.T) {
	t.Run("should return empty string when context does not have flatten key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		result := GetCNAMEFlatten(ctx)
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	t.Run("should return value when context has flatten key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "flattenValue"
		ctx := context.WithValue(context.Background(), flattenKey, expected)
		result := GetCNAMEFlatten(ctx)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
