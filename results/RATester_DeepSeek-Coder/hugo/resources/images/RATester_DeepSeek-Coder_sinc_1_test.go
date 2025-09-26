package images

import (
	"fmt"
	"testing"
)

func TestSinc_1(t *testing.T) {
	t.Run("Testing sinc function with positive values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := sinc(float32(0.5))
		expected := float32(0.63661977)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Testing sinc function with zero", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := sinc(float32(0))
		expected := float32(1)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Testing sinc function with negative values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := sinc(float32(-0.5))
		expected := float32(0.63661977)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})
}
