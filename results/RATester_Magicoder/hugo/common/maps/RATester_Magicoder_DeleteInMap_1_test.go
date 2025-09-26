package maps

import (
	"fmt"
	"testing"
)

func TestDeleteInMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: map[string]any{
			"key1": map[string]any{
				"mapKey1": "value1",
				"mapKey2": "value2",
			},
		},
	}

	scratch.DeleteInMap("key1", "mapKey1")

	if _, ok := scratch.values["key1"].(map[string]any)["mapKey1"]; ok {
		t.Error("Expected mapKey1 to be deleted")
	}
}
