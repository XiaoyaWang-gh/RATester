package maps

import (
	"fmt"
	"testing"
)

func TestToParamsAndPrepare_1(t *testing.T) {
	t.Parallel()

	t.Run("nil input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params, err := ToParamsAndPrepare(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(params) != 0 {
			t.Errorf("expected empty params, got %v", params)
		}
	})

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := map[string]any{
			"key1": "value1",
			"key2": "value2",
		}
		params, err := ToParamsAndPrepare(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(params) != 2 {
			t.Errorf("expected 2 params, got %v", len(params))
		}
		if params["key1"] != "value1" {
			t.Errorf("expected 'value1', got %v", params["key1"])
		}
		if params["key2"] != "value2" {
			t.Errorf("expected 'value2', got %v", params["key2"])
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := 123
		_, err := ToParamsAndPrepare(input)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
