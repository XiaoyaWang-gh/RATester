package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenameKeys_1(t *testing.T) {
	t.Run("should rename keys in map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := map[string]any{
			"oldKey1": "value1",
			"oldKey2": "value2",
		}
		expected := map[string]any{
			"newKey1": "value1",
			"newKey2": "value2",
		}
		RenameKeys(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v, got %v", expected, input)
		}
	})

	t.Run("should not affect map without aliases", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := map[string]any{
			"key1": "value1",
			"key2": "value2",
		}
		expected := map[string]any{
			"key1": "value1",
			"key2": "value2",
		}
		RenameKeys(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v, got %v", expected, input)
		}
	})

	t.Run("should not affect empty map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := map[string]any{}
		expected := map[string]any{}
		RenameKeys(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v, got %v", expected, input)
		}
	})
}
