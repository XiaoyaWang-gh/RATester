package cache

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	t.Run("Adding a new entry", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(10)
		cache.Add("key1", "value1")
		val, ok := cache.Get("key1")
		if !ok || val != "value1" {
			t.Errorf("Expected value 'value1' for key 'key1', got %v", val)
		}
	})

	t.Run("Adding an existing entry", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(10)
		cache.Add("key1", "value1")
		cache.Add("key1", "value2")
		val, ok := cache.Get("key1")
		if !ok || val != "value2" {
			t.Errorf("Expected value 'value2' for key 'key1', got %v", val)
		}
	})

	t.Run("Adding an entry with limit", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(1)
		cache.Add("key1", "value1")
		cache.Add("key2", "value2")
		_, ok := cache.Get("key1")
		if ok {
			t.Errorf("Expected key 'key1' to be evicted")
		}
		val, ok := cache.Get("key2")
		if !ok || val != "value2" {
			t.Errorf("Expected value 'value2' for key 'key2', got %v", val)
		}
	})
}
