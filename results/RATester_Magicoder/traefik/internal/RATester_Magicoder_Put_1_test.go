package main

import (
	"fmt"
	"testing"
)

func TestPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sw := storeWriter{
		data: make(map[string]string),
	}

	key := "testKey"
	value := []byte("testValue")

	err := sw.Put(key, value, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	storedValue, ok := sw.data[key]
	if !ok {
		t.Errorf("Expected key %s to be in the store", key)
	}

	if storedValue != string(value) {
		t.Errorf("Expected value %s, got %s", value, storedValue)
	}
}
