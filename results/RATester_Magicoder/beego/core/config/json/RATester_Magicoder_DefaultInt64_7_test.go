package json

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": int64(10),
			"key2": "not_an_int",
		},
	}

	// Test case 1: key exists, value is int64
	val := container.DefaultInt64("key1", 20)
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// Test case 2: key exists, value is not int64
	val = container.DefaultInt64("key2", 20)
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	// Test case 3: key does not exist
	val = container.DefaultInt64("key3", 30)
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
}
