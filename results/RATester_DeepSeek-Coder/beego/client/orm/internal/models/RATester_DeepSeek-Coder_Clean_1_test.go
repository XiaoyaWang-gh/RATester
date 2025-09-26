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

	t.Parallel()

	mc := &ModelCache{
		orders:          []string{"order1", "order2"},
		cache:           map[string]*ModelInfo{"cache1": {Name: "name1"}, "cache2": {Name: "name2"}},
		cacheByFullName: map[string]*ModelInfo{"fullname1": {Name: "name1"}, "fullname2": {Name: "name2"}},
		done:            true,
	}

	mc.Clean()

	if len(mc.orders) != 0 {
		t.Errorf("Expected orders to be empty, got %v", mc.orders)
	}

	if len(mc.cache) != 0 {
		t.Errorf("Expected cache to be empty, got %v", mc.cache)
	}

	if len(mc.cacheByFullName) != 0 {
		t.Errorf("Expected cacheByFullName to be empty, got %v", mc.cacheByFullName)
	}

	if mc.done != false {
		t.Errorf("Expected done to be false, got %v", mc.done)
	}
}
