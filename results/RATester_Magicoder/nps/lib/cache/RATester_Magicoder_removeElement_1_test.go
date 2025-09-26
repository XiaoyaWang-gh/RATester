package cache

import (
	"container/list"
	"fmt"
	"testing"
)

func TestremoveElement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		MaxEntries: 10,
		ll:         list.New(),
	}
	key := "testKey"
	value := "testValue"
	cache.cache.Store(key, value)
	element := cache.ll.PushFront(&entry{key: key, value: value})
	cache.removeElement(element)
	if _, ok := cache.cache.Load(key); ok {
		t.Errorf("Expected key to be removed from cache")
	}
	if cache.ll.Len() != 0 {
		t.Errorf("Expected linked list to be empty")
	}
}
