package json

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": float64(1.23),
			"key2": "not a float",
		},
	}

	// Test with a valid float
	result := container.DefaultFloat("key1", 0.0)
	if result != 1.23 {
		t.Errorf("Expected 1.23, got %f", result)
	}

	// Test with an invalid float
	result = container.DefaultFloat("key2", 4.56)
	if result != 4.56 {
		t.Errorf("Expected 4.56, got %f", result)
	}
}
