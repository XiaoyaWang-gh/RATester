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

	c := &Scratch{
		values: map[string]any{},
	}
	key := "key"
	newAddend := 1
	_, err := c.Add(key, newAddend)
	if err != nil {
		t.Errorf("Add() error = %v", err)
		return
	}
	if c.values[key] != newAddend {
		t.Errorf("Add() = %v, want %v", c.values[key], newAddend)
	}
}
