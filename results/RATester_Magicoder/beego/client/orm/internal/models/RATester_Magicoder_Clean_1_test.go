package models

import (
	"fmt"
	"testing"
)

func TestClean_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		orders:          []string{"test1", "test2"},
		cache:           map[string]*ModelInfo{"test1": &ModelInfo{}, "test2": &ModelInfo{}},
		cacheByFullName: map[string]*ModelInfo{"test1": &ModelInfo{}, "test2": &ModelInfo{}},
		done:            true,
	}

	mc.Clean()

	if len(mc.orders) != 0 {
		t.Errorf("Expected orders to be empty, but got %v", mc.orders)
	}

	if len(mc.cache) != 0 {
		t.Errorf("Expected cache to be empty, but got %v", mc.cache)
	}

	if len(mc.cacheByFullName) != 0 {
		t.Errorf("Expected cacheByFullName to be empty, but got %v", mc.cacheByFullName)
	}

	if mc.done != false {
		t.Errorf("Expected done to be false, but got %v", mc.done)
	}
}
