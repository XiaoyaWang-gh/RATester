package maps

import (
	"fmt"
	"testing"
)

func TestAdd_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: make(map[string]any),
	}

	key := "testKey"
	newAddend := 10

	_, err := scratch.Add(key, newAddend)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, found := scratch.values[key]
	if !found {
		t.Errorf("Expected key %s not found in scratch", key)
	}

	if val != newAddend {
		t.Errorf("Expected value %v, got %v", newAddend, val)
	}
}
