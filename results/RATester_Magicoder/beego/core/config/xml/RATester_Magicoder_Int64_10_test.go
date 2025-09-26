package xml

import (
	"fmt"
	"testing"
)

func TestInt64_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &ConfigContainer{
		data: map[string]interface{}{
			"key": "123",
		},
	}

	val, err := container.Int64("key")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if val != 123 {
		t.Errorf("Expected 123, but got %v", val)
	}
}
