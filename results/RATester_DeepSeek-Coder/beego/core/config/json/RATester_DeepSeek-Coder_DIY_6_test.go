package json

import (
	"fmt"
	"testing"
)

func TestDIY_6(t *testing.T) {
	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
	}

	t.Run("exist key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val, err := container.DIY("key1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if val != "value1" {
			t.Errorf("expected value 'value1', got %v", val)
		}
	})

	t.Run("not exist key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		val, err := container.DIY("key4")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if val != nil {
			t.Errorf("expected nil, got %v", val)
		}
	})
}
