package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestItems_1(t *testing.T) {
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
	b.Set("key2", "value2")
	b.Set("key3", "value3")

	items := b.Items()

	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	if items["key1"] != "value1" {
		t.Errorf("Expected 'value1', got '%v'", items["key1"])
	}

	if items["key2"] != "value2" {
		t.Errorf("Expected 'value2', got '%v'", items["key2"])
	}

	if items["key3"] != "value3" {
		t.Errorf("Expected 'value3', got '%v'", items["key3"])
	}
}
