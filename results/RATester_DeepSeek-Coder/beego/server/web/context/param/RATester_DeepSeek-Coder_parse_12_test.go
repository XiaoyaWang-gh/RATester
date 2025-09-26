package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse_12(t *testing.T) {
	t.Run("Test float32 parse", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		parser := floatParser{}
		value := "3.14"
		toType := reflect.TypeOf(float32(0))
		expected := float32(3.14)
		result, err := parser.parse(value, toType)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test float64 parse", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		parser := floatParser{}
		value := "3.14"
		toType := reflect.TypeOf(float64(0))
		expected := float64(3.14)
		result, err := parser.parse(value, toType)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test parse error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		parser := floatParser{}
		value := "invalid"
		toType := reflect.TypeOf(float64(0))
		_, err := parser.parse(value, toType)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
