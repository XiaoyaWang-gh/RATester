package models

import (
	"fmt"
	"testing"
)

func TestSetRaw_5(t *testing.T) {
	e := new(IntegerField)

	t.Run("Test with int32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw(int32(123))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if e.Value() != 123 {
			t.Errorf("Expected 123, got %v", e.Value())
		}
	})

	t.Run("Test with string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw("456")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if e.Value() != 456 {
			t.Errorf("Expected 456, got %v", e.Value())
		}
	})

	t.Run("Test with unknown type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := e.SetRaw([]byte("789"))
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if e.Value() != 456 {
			t.Errorf("Expected 456, got %v", e.Value())
		}
	})
}
