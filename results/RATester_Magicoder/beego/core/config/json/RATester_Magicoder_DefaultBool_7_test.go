package json

import (
	"fmt"
	"testing"
)

func TestDefaultBool_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": true,
			"key2": false,
		},
	}

	// Test case 1: key exists, return value
	val := container.DefaultBool("key1", false)
	if val != true {
		t.Errorf("Expected true, got %v", val)
	}

	// Test case 2: key does not exist, return default value
	val = container.DefaultBool("key3", true)
	if val != true {
		t.Errorf("Expected true, got %v", val)
	}
}
