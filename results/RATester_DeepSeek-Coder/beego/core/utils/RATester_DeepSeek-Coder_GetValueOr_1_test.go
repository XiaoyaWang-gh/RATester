package utils

import (
	"fmt"
	"testing"
)

func TestGetValueOr_1(t *testing.T) {
	kvs := &SimpleKVs{kvs: map[interface{}]interface{}{"key1": "value1", "key2": "value2"}}

	t.Run("TestGetValueOr_ExistingKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "value1"
		actual := kvs.GetValueOr("key1", "default")
		if actual != expected {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})

	t.Run("TestGetValueOr_NonExistingKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "default"
		actual := kvs.GetValueOr("non-existing-key", "default")
		if actual != expected {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})
}
