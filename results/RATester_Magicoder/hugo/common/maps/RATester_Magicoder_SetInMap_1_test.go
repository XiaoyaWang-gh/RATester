package maps

import (
	"fmt"
	"testing"
)

func TestSetInMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: make(map[string]any),
	}

	key := "testKey"
	mapKey := "testMapKey"
	value := "testValue"

	scratch.SetInMap(key, mapKey, value)

	val, ok := scratch.values[key].(map[string]any)[mapKey]
	if !ok {
		t.Errorf("Expected value to be set in map")
	}

	if val != value {
		t.Errorf("Expected value to be %v, but got %v", value, val)
	}
}
