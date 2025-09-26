package echo

import (
	"fmt"
	"testing"
)

func TestMustFloat32_1(t *testing.T) {
	t.Run("Test MustFloat32 with valid float32 value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				if sourceParam == "validFloat32" {
					return "1.23"
				}
				return ""
			},
		}
		var dest float32
		b.MustFloat32("validFloat32", &dest)
		if dest != float32(1.23) {
			t.Errorf("Expected dest to be 1.23, got %v", dest)
		}
	})

	t.Run("Test MustFloat32 with invalid float32 value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				if sourceParam == "invalidFloat32" {
					return "invalid"
				}
				return ""
			},
		}
		var dest float32
		err := b.MustFloat32("invalidFloat32", &dest)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Test MustFloat32 with non-existing value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValueFunc: func(sourceParam string) string {
				return ""
			},
		}
		var dest float32
		err := b.MustFloat32("nonExisting", &dest)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
