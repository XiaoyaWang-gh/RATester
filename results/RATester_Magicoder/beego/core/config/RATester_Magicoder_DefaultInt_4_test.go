package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {"key1": "10", "key2": "20"},
		},
	}

	// Test case 1: key exists, return the value
	val := container.DefaultInt("section1:key1", 0)
	if val != 10 {
		t.Errorf("Expected 10, but got %d", val)
	}

	// Test case 2: key does not exist, return default value
	val = container.DefaultInt("section1:key3", 30)
	if val != 30 {
		t.Errorf("Expected 30, but got %d", val)
	}

	// Test case 3: section does not exist, return default value
	val = container.DefaultInt("section2:key1", 40)
	if val != 40 {
		t.Errorf("Expected 40, but got %d", val)
	}
}
