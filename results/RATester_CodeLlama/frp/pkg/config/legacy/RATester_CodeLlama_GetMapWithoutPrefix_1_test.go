package legacy

import (
	"fmt"
	"testing"
)

func TestGetMapWithoutPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	set := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	m := GetMapWithoutPrefix(set, "a")

	if len(m) != 2 {
		t.Errorf("Expected 2, got %d", len(m))
	}

	if m["b"] != "2" {
		t.Errorf("Expected 2, got %s", m["b"])
	}

	if m["c"] != "3" {
		t.Errorf("Expected 3, got %s", m["c"])
	}
}
