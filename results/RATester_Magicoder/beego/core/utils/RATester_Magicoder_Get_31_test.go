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

	bm := &BeeMap{
		lock: &sync.RWMutex{},
		bm:   make(map[interface{}]interface{}),
	}

	key := "testKey"
	value := "testValue"

	bm.Set(key, value)

	got := bm.Get(key)
	if got != value {
		t.Errorf("Get() = %v, want %v", got, value)
	}

	got = bm.Get("nonExistentKey")
	if got != nil {
		t.Errorf("Get() = %v, want %v", got, nil)
	}
}
