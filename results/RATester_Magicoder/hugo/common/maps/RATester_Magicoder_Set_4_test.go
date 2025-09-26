package maps

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: make(map[string]any),
	}

	key := "testKey"
	value := "testValue"

	scratch.Set(key, value)

	if scratch.values[key] != value {
		t.Errorf("Expected %v, got %v", value, scratch.values[key])
	}
}
