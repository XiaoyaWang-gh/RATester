package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetMapKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testMap := sync.Map{}
	testMap.Store(1, "one")
	testMap.Store(2, "two")
	testMap.Store(3, "three")

	keys := GetMapKeys(testMap, false, "", "")

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	for _, key := range keys {
		if _, ok := testMap.Load(key); !ok {
			t.Errorf("Key %d not found in map", key)
		}
	}
}
