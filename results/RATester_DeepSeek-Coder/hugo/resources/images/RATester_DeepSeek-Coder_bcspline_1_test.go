package images

import (
	"fmt"
	"testing"
)

func TestBcspline_1(t *testing.T) {
	t.Run("Test with x=0, b=0, c=0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		x := float32(0)
		b := float32(0)
		c := float32(0)
		expected := float32(0)
		result := bcspline(x, b, c)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Test with x=1, b=1, c=1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		x := float32(1)
		b := float32(1)
		c := float32(1)
		expected := float32(0.5)
		result := bcspline(x, b, c)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Test with x=2, b=2, c=2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		x := float32(2)
		b := float32(2)
		c := float32(2)
		expected := float32(1)
		result := bcspline(x, b, c)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Test with x=-1, b=-1, c=-1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		x := float32(-1)
		b := float32(-1)
		c := float32(-1)
		expected := float32(0.5)
		result := bcspline(x, b, c)
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})
}
