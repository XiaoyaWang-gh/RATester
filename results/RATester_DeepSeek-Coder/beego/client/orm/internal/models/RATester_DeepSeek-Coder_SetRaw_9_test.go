package models

import (
	"fmt"
	"testing"
)

func TestSetRaw_9(t *testing.T) {
	e := new(BigIntegerField)

	t.Run("Test with int64 value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw(int64(1234567890))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if e.Value() != 1234567890 {
			t.Errorf("Expected value to be 1234567890, got %v", e.Value())
		}
	})

	t.Run("Test with string value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw("9876543210")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if e.Value() != 9876543210 {
			t.Errorf("Expected value to be 9876543210, got %v", e.Value())
		}
	})

	t.Run("Test with invalid string value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw("invalid")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if e.Value() != 9876543210 {
			t.Errorf("Expected value to be 9876543210, got %v", e.Value())
		}
	})

	t.Run("Test with unsupported type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw([]int{1, 2, 3})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if e.Value() != 9876543210 {
			t.Errorf("Expected value to be 9876543210, got %v", e.Value())
		}
	})
}
