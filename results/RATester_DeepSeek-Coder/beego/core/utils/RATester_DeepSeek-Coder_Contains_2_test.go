package utils

import (
	"fmt"
	"testing"
)

func TestContains_2(t *testing.T) {
	kvs := &SimpleKVs{kvs: map[interface{}]interface{}{"key1": "value1", "key2": "value2"}}

	t.Run("Contains existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if !kvs.Contains("key1") {
			t.Errorf("Expected key1 to be in the map")
		}
	})

	t.Run("Does not Contains non-existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if kvs.Contains("key3") {
			t.Errorf("Expected key3 not to be in the map")
		}
	})
}
