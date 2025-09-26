package math

import (
	"fmt"
	"testing"
)

func TestMul_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("TestMul_PositiveNumbers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Mul(2, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 6 {
			t.Errorf("Expected 6, got %v", result)
		}
	})

	t.Run("TestMul_NegativeNumbers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Mul(-2, -3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 6 {
			t.Errorf("Expected 6, got %v", result)
		}
	})

	t.Run("TestMul_PositiveAndNegativeNumbers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Mul(-2, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != -6 {
			t.Errorf("Expected -6, got %v", result)
		}
	})

	t.Run("TestMul_Zero", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Mul(0, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected 0, got %v", result)
		}
	})

	t.Run("TestMul_MultipleNumbers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Mul(2, 3, 4)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 24 {
			t.Errorf("Expected 24, got %v", result)
		}
	})
}
