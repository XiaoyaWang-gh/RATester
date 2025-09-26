package client

import (
	"fmt"
	"testing"
)

func TestVisitAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := make(PathParam)
	pp.SetParam("key1", "val1")
	pp.SetParam("key2", "val2")

	visited := make(map[string]string)
	pp.VisitAll(func(key, val string) {
		visited[key] = val
	})

	if len(visited) != 2 {
		t.Errorf("Expected 2 items, got %d", len(visited))
	}

	if visited["key1"] != "val1" {
		t.Errorf("Expected 'val1', got '%s'", visited["key1"])
	}

	if visited["key2"] != "val2" {
		t.Errorf("Expected 'val2', got '%s'", visited["key2"])
	}
}
