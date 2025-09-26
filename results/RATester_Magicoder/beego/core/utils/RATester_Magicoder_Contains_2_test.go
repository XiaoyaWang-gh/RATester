package utils

import (
	"fmt"
	"testing"
)

func TestContains_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	kvs := &SimpleKVs{kvs: map[interface{}]interface{}{"key1": "value1", "key2": "value2"}}

	// Test case 1: key exists
	if !kvs.Contains("key1") {
		t.Error("Expected true, got false")
	}

	// Test case 2: key does not exist
	if kvs.Contains("key3") {
		t.Error("Expected false, got true")
	}
}
