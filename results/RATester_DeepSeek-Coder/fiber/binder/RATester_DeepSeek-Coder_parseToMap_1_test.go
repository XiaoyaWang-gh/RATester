package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseToMap_1(t *testing.T) {
	t.Run("Test with map[string]string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		data := map[string][]string{
			"key1": {"value1", "value2"},
			"key2": {"value3", "value4"},
		}
		expected := map[string]string{
			"key1": "value2",
			"key2": "value4",
		}
		var result map[string]string
		err := parseToMap(&result, data)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test with map[string][]string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		data := map[string][]string{
			"key1": {"value1", "value2"},
			"key2": {"value3", "value4"},
		}
		expected := map[string][]string{
			"key1": {"value1", "value2"},
			"key2": {"value3", "value4"},
		}
		var result map[string][]string
		err := parseToMap(&result, data)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test with invalid type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		data := map[string][]string{
			"key1": {"value1", "value2"},
			"key2": {"value3", "value4"},
		}
		var result int
		err := parseToMap(&result, data)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
