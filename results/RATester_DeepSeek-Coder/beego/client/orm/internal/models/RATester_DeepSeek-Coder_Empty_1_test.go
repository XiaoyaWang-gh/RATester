package models

import (
	"fmt"
	"testing"
)

func TestEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	mc := &ModelCache{
		cache: make(map[string]*ModelInfo),
	}

	if !mc.Empty() {
		t.Errorf("Expected cache to be empty")
	}

	mc.cache["test"] = &ModelInfo{}

	if mc.Empty() {
		t.Errorf("Expected cache not to be empty")
	}
}
