package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestSet_41(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bm := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	// Test case 1: Set a new key-value pair
	key1 := "key1"
	value1 := "value1"
	expected1 := true
	result1 := bm.Set(key1, value1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Set an existing key with a different value
	key2 := "key2"
	value2 := "value2"
	value2_new := "value2_new"
	expected2 := true
	bm.Set(key2, value2)
	result2 := bm.Set(key2, value2_new)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Set an existing key with the same value
	key3 := "key3"
	value3 := "value3"
	expected3 := false
	bm.Set(key3, value3)
	result3 := bm.Set(key3, value3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}
}
