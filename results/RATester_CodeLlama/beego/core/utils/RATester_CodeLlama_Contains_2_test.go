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

	kvs := &SimpleKVs{
		kvs: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}
	key := "key1"
	if !kvs.Contains(key) {
		t.Errorf("key1 should be in kvs")
	}
	key = "key3"
	if kvs.Contains(key) {
		t.Errorf("key3 should not be in kvs")
	}
}
