package cache

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := New(2)
	cache.Add("key1", "value1")
	cache.Add("key2", "value2")

	value, ok := cache.Get("key1")
	if !ok {
		t.Fatal("Expected key1 to exist")
	}
	if value != "value1" {
		t.Fatalf("Expected value1, got %v", value)
	}

	value, ok = cache.Get("key3")
	if ok {
		t.Fatal("Expected key3 to not exist")
	}
	if value != nil {
		t.Fatalf("Expected nil, got %v", value)
	}
}
