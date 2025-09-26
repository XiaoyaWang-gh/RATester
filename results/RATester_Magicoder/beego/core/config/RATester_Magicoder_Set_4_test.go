package config

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &fakeConfigContainer{
		data: make(map[string]string),
	}

	err := fc.Set("Key", "Value")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fc.data["key"] != "Value" {
		t.Errorf("Expected 'Value', got '%v'", fc.data["key"])
	}
}
