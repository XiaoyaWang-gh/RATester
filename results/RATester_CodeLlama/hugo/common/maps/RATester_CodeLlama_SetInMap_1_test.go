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

	c := &Scratch{}
	key := "key"
	mapKey := "mapKey"
	value := "value"
	c.SetInMap(key, mapKey, value)
	if c.values[key].(map[string]any)[mapKey] != value {
		t.Errorf("SetInMap failed")
	}
}
