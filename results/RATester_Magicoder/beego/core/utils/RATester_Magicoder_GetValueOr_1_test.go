package utils

import (
	"fmt"
	"testing"
)

func TestGetValueOr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	kvs := &SimpleKVs{kvs: map[interface{}]interface{}{"key1": "value1", "key2": "value2"}}

	// Test case 1: key exists
	value := kvs.GetValueOr("key1", "default")
	if value != "value1" {
		t.Errorf("Expected 'value1', but got '%v'", value)
	}

	// Test case 2: key does not exist
	value = kvs.GetValueOr("key3", "default")
	if value != "default" {
		t.Errorf("Expected 'default', but got '%v'", value)
	}
}
