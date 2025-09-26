package json

import (
	"fmt"
	"testing"
)

func TestDIY_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Test case 1: key exists
	val, err := container.DIY("key1")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected value 'value1', but got %v", val)
	}

	// Test case 2: key does not exist
	val, err = container.DIY("key3")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if val != nil {
		t.Errorf("Expected nil, but got %v", val)
	}
}
