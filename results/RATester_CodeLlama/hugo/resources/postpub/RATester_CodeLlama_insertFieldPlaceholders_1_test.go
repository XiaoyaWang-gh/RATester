package postpub

import (
	"fmt"
	"testing"
)

func TestInsertFieldPlaceholders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	root := "root"
	m := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}
	createPlaceholder := func(s string) string {
		return "placeholder"
	}
	insertFieldPlaceholders(root, m, createPlaceholder)
	if m["key1"] != "placeholder" {
		t.Errorf("m[\"key1\"] = %v, want %v", m["key1"], "placeholder")
	}
	if m["key2"] != "placeholder" {
		t.Errorf("m[\"key2\"] = %v, want %v", m["key2"], "placeholder")
	}
}
