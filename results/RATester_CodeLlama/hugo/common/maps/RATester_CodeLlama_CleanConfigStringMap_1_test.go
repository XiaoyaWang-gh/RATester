package maps

import (
	"fmt"
	"testing"
)

func TestCleanConfigStringMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := map[string]any{
		"a": "b",
		"c": "d",
		"e": "f",
	}
	m2 := CleanConfigStringMap(m)
	if len(m2) != 3 {
		t.Errorf("m2 length is %d, want 3", len(m2))
	}
	if m2["a"] != "b" {
		t.Errorf("m2[\"a\"] is %s, want b", m2["a"])
	}
	if m2["c"] != "d" {
		t.Errorf("m2[\"c\"] is %s, want d", m2["c"])
	}
	if m2["e"] != "f" {
		t.Errorf("m2[\"e\"] is %s, want f", m2["e"])
	}
}
