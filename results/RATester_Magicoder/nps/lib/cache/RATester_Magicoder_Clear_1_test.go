package cache

import (
	"fmt"
	"testing"
)

func TestClear_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := New(10)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Clear()
	if cache.Len() != 0 {
		t.Errorf("Expected cache length to be 0, but got %d", cache.Len())
	}
}
