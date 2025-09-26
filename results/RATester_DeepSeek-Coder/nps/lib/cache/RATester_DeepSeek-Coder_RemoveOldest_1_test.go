package cache

import (
	"fmt"
	"testing"
)

func TestRemoveOldest_1(t *testing.T) {
	t.Run("RemoveOldest", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cache := New(2)
		cache.Add("1", 1)
		cache.Add("2", 2)
		cache.RemoveOldest()
		if _, ok := cache.Get("1"); ok {
			t.Errorf("Expected to not find key '1'")
		}
		if _, ok := cache.Get("2"); !ok {
			t.Errorf("Expected to find key '2'")
		}
	})
}
