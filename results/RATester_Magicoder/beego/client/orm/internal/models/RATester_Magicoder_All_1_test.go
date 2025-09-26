package models

import (
	"fmt"
	"testing"
)

func TestAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache: map[string]*ModelInfo{
			"test1": {Name: "test1"},
			"test2": {Name: "test2"},
		},
	}

	result := mc.All()

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result["test1"].Name != "test1" {
		t.Errorf("Expected test1, got %s", result["test1"].Name)
	}

	if result["test2"].Name != "test2" {
		t.Errorf("Expected test2, got %s", result["test2"].Name)
	}
}
