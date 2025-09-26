package memory

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

	storage := &Storage{
		data: map[string]item{
			"key1": {v: "value1", e: 0},
			"key2": {v: "value2", e: 0},
		},
	}

	storage.Delete("key1")

	if _, ok := storage.data["key1"]; ok {
		t.Error("Expected key1 to be deleted")
	}

	if _, ok := storage.data["key2"]; !ok {
		t.Error("Expected key2 to still exist")
	}
}
