package client

import (
	"fmt"
	"testing"
)

func TestDelParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := PathParam{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}

	pp.DelParams("key1", "key2")

	if len(pp) != 1 {
		t.Errorf("Expected length of PathParam to be 1, but got %d", len(pp))
	}

	if _, ok := pp["key1"]; ok {
		t.Errorf("Expected key1 to be deleted, but it still exists")
	}

	if _, ok := pp["key2"]; ok {
		t.Errorf("Expected key2 to be deleted, but it still exists")
	}

	if _, ok := pp["key3"]; !ok {
		t.Errorf("Expected key3 to still exist, but it was deleted")
	}
}
