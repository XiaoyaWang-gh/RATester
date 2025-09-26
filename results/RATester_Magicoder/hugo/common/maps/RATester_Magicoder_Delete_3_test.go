package maps

import (
	"fmt"
	"testing"
)

func TestDelete_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	scratch.Delete("key1")

	if _, ok := scratch.values["key1"]; ok {
		t.Error("Expected key1 to be deleted")
	}
}
