package math

import (
	"fmt"
	"testing"
)

func TestSub_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("Subtracts integers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sub(10, 5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 5 {
			t.Errorf("Expected 5, got %v", result)
		}
	})

	t.Run("Subtracts floats", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sub(10.5, 5.2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 5.3 {
			t.Errorf("Expected 5.3, got %v", result)
		}
	})

	t.Run("Subtracts integers and floats", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sub(10, 5.2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 4.8 {
			t.Errorf("Expected 4.8, got %v", result)
		}
	})

	t.Run("Subtracts negative numbers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sub(10, -5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 15 {
			t.Errorf("Expected 15, got %v", result)
		}
	})

	t.Run("Returns error for invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Sub("invalid")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
