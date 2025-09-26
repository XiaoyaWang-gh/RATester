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

	bm := &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}

	bm.Set("key1", "value1")
	bm.Set("key2", "value2")

	items := bm.Items()

	if len(items) != 2 {
		t.Error("Items() failed. Expected 2, got ", len(items))
	}

	if items["key1"] != "value1" {
		t.Error("Items() failed. Expected 'value1', got ", items["key1"])
	}

	if items["key2"] != "value2" {
		t.Error("Items() failed. Expected 'value2', got ", items["key2"])
	}
}
