package config

import (
	"fmt"
	"testing"
)

func TestRenameKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	keyAliases.Rename(m)
	if len(m) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(m))
	}
	if m["key1"] != "value1" {
		t.Errorf("Expected key1 to be value1, got %s", m["key1"])
	}
	if m["key2"] != "value2" {
		t.Errorf("Expected key2 to be value2, got %s", m["key2"])
	}
	if m["key3"] != "value3" {
		t.Errorf("Expected key3 to be value3, got %s", m["key3"])
	}
}
