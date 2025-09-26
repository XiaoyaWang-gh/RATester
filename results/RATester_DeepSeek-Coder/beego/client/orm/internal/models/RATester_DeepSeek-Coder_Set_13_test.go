package models

import (
	"fmt"
	"testing"
)

func TestSet_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache:           make(map[string]*ModelInfo),
		cacheByFullName: make(map[string]*ModelInfo),
	}

	mi := &ModelInfo{
		FullName: "test",
	}

	oldMI := mc.Set("test", mi)

	if oldMI != nil {
		t.Errorf("Expected nil, got %v", oldMI)
	}

	if mc.cache["test"] != mi {
		t.Errorf("Expected %v, got %v", mi, mc.cache["test"])
	}

	if mc.cacheByFullName["test"] != mi {
		t.Errorf("Expected %v, got %v", mi, mc.cacheByFullName["test"])
	}

	if len(mc.orders) != 1 || mc.orders[0] != "test" {
		t.Errorf("Expected [\"test\"], got %v", mc.orders)
	}
}
