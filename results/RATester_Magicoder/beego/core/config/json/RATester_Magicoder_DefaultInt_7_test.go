package json

import (
	"fmt"
	"testing"
)

func TestDefaultInt_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": 10,
			"key2": "not an int",
		},
	}

	// Test with a valid key
	result := container.DefaultInt("key1", 20)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	// Test with an invalid key
	result = container.DefaultInt("key2", 20)
	if result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}

	// Test with a non-existent key
	result = container.DefaultInt("key3", 30)
	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}
