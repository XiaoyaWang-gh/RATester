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

	s := &Storage{
		data: map[string]item{
			"key1": {v: "value1", e: 100},
			"key2": {v: "value2", e: 200},
		},
	}

	s.Reset()

	if len(s.data) != 0 {
		t.Errorf("Expected data to be empty after reset, but got %d items", len(s.data))
	}
}
