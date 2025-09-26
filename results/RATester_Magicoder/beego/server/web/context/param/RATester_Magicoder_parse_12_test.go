package param

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_12(t *testing.T) {
	parser := floatParser{}

	t.Run("Test parse float32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "3.14"
		toType := reflect.TypeOf(float32(0))
		expected := float32(3.14)

		result, err := parser.parse(value, toType)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test parse float64", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "3.14"
		toType := reflect.TypeOf(float64(0))
		expected := float64(3.14)

		result, err := parser.parse(value, toType)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test parse error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "invalid"
		toType := reflect.TypeOf(float64(0))

		_, err := parser.parse(value, toType)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
