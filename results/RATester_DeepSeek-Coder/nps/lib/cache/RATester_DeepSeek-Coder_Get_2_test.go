package cache

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	t.Run("should return false if key does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(0)
		_, ok := cache.Get("non-existent-key")
		if ok {
			t.Errorf("Expected ok to be false")
		}
	})

	t.Run("should return true and value if key exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(0)
		cache.Add("existing-key", "existing-value")
		value, ok := cache.Get("existing-key")
		if !ok {
			t.Errorf("Expected ok to be true")
		}
		if value != "existing-value" {
			t.Errorf("Expected value to be 'existing-value'")
		}
	})

	t.Run("should move accessed key to the front of the list", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(2)
		cache.Add("key1", "value1")
		cache.Add("key2", "value2")
		cache.Get("key1")
		front := cache.ll.Front()
		if front.Value.(*entry).key != "key1" {
			t.Errorf("Expected key at front of list to be 'key1'")
		}
	})
}
