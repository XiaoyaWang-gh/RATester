package config

import (
	"fmt"
	"testing"
)

func TestDIY_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {"key1": "value1", "key2": "value2"},
			"section2": {"key3": "value3", "key4": "value4"},
		},
	}

	// Test case 1: key exists
	v, err := container.DIY("section1:key1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if v != "value1" {
		t.Errorf("Expected value1, got %v", v)
	}

	// Test case 2: key does not exist
	_, err = container.DIY("section1:key3")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test case 3: section does not exist
	_, err = container.DIY("section3:key1")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
