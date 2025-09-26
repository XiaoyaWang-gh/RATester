package cache

import (
	"fmt"
	"testing"
)

func TestRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := New(10)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Remove("key1")
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("Expected key1 to be removed")
	}
}
