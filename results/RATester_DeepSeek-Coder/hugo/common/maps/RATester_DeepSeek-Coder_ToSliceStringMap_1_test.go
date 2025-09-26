package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToSliceStringMap_1(t *testing.T) {
	t.Run("Testing with []map[string]any", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := []map[string]any{{"key": "value"}}
		output, err := ToSliceStringMap(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(output, input) {
			t.Errorf("Expected %v, got %v", input, output)
		}
	})

	t.Run("Testing with Params", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := Params{"key": "value"}
		output, err := ToSliceStringMap(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(output, []map[string]any{input}) {
			t.Errorf("Expected %v, got %v", []map[string]any{input}, output)
		}
	})

	t.Run("Testing with []any", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := []any{map[string]any{"key": "value"}}
		output, err := ToSliceStringMap(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(output, []map[string]any{{"key": "value"}}) {
			t.Errorf("Expected %v, got %v", []map[string]any{{"key": "value"}}, output)
		}
	})

	t.Run("Testing with unsupported type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "unsupported"
		_, err := ToSliceStringMap(input)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
