package json

import (
	"fmt"
	"testing"
)

func TestInt64_7(t *testing.T) {
	t.Run("TestInt64_ExistKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		container := &JSONConfigContainer{
			data: map[string]interface{}{
				"key": float64(1234567890),
			},
		}
		val, err := container.Int64("key")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if val != 1234567890 {
			t.Errorf("Expected 1234567890, got %v", val)
		}
	})

	t.Run("TestInt64_NotExistKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		container := &JSONConfigContainer{
			data: map[string]interface{}{},
		}
		_, err := container.Int64("key")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("TestInt64_NotInt64Value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		container := &JSONConfigContainer{
			data: map[string]interface{}{
				"key": "not int64",
			},
		}
		_, err := container.Int64("key")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
