package echo

import (
	"fmt"
	"testing"
)

func TestFloat_1(t *testing.T) {
	t.Parallel()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			switch sourceParam {
			case "validFloat64":
				return "123.456"
			case "validFloat32":
				return "123.456"
			case "invalidFloat":
				return "not a number"
			default:
				return ""
			}
		},
		ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
			return fmt.Errorf("sourceParam: %s, values: %v, message: %v, internalError: %v", sourceParam, values, message, internalError)
		},
	}

	t.Run("validFloat64", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest float64
		b.float("validFloat64", b.ValueFunc("validFloat64"), &dest, 64)
		if dest != 123.456 {
			t.Errorf("Expected 123.456, got %v", dest)
		}
	})

	t.Run("validFloat32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest float32
		b.float("validFloat32", b.ValueFunc("validFloat32"), &dest, 32)
		if dest != 123.456 {
			t.Errorf("Expected 123.456, got %v", dest)
		}
	})

	t.Run("invalidFloat", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest float64
		b.float("invalidFloat", b.ValueFunc("invalidFloat"), &dest, 64)
		if dest != 0 {
			t.Errorf("Expected 0, got %v", dest)
		}
	})
}
