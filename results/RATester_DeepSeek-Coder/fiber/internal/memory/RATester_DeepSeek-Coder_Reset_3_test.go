package memory

import (
	"fmt"
	"testing"
)

func TestReset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &Storage{
		data: map[string]item{
			"key1": {v: "value1", e: 123456789},
			"key2": {v: "value2", e: 987654321},
		},
	}

	store.Reset()

	if len(store.data) != 0 {
		t.Errorf("Expected data to be empty after reset, but got %d items", len(store.data))
	}
}
