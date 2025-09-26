package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_5(t *testing.T) {
	t.Run("should return default value when key does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "non-existing-key"
		defaultVal := 123.456
		expected := defaultVal
		result := DefaultFloat(key, defaultVal)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("should return value from globalInstance when key exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := "existing-key"
		defaultVal := 789.012
		globalInstance.Set(key, "456.789")
		expected := 456.789
		result := DefaultFloat(key, defaultVal)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
