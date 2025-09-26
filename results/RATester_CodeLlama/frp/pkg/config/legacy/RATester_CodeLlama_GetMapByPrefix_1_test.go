package legacy

import (
	"fmt"
	"testing"
)

func TestGetMapByPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	set := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
	}

	m := GetMapByPrefix(set, "a")

	if len(m) != 1 {
		t.Errorf("Expected 1, got %d", len(m))
	}

	if m["a"] != "1" {
		t.Errorf("Expected 1, got %s", m["a"])
	}
}
