package cache

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetFloat64_1(t *testing.T) {
	t.Run("Test float64", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := float64(123.456)
		result := GetFloat64(value)
		if result != value {
			t.Errorf("Expected %v, got %v", value, result)
		}
	})

	t.Run("Test string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "789.123"
		result := GetFloat64(value)
		expected, _ := strconv.ParseFloat(value, 64)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test other types", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "not a number"
		result := GetFloat64(value)
		if result != 0 {
			t.Errorf("Expected 0, got %v", result)
		}
	})
}
