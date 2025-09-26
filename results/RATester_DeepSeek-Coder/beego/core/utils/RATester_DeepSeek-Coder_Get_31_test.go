package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestGet_31(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	b.Set("key1", "value1")

	val := b.Get("key1")
	if val != "value1" {
		t.Errorf("Expected 'value1', got '%v'", val)
	}

	val = b.Get("key2")
	if val != nil {
		t.Errorf("Expected nil, got '%v'", val)
	}
}
