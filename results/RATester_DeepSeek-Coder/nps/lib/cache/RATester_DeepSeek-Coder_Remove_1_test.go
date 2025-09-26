package cache

import (
	"container/list"
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
	t.Run("Remove existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := &Cache{
			MaxEntries: 5,
		}
		c.cache.Store("key1", &list.Element{})
		c.Remove("key1")
		_, ok := c.cache.Load("key1")
		if ok {
			t.Errorf("Expected key to be removed")
		}
	})

	t.Run("Remove non-existing key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := &Cache{
			MaxEntries: 5,
		}
		c.Remove("key2")
		_, ok := c.cache.Load("key2")
		if ok {
			t.Errorf("Expected key to not exist")
		}
	})
}
