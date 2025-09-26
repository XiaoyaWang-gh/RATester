package collections

import (
	"fmt"
	"testing"
)

func TestNewSliceElement_1(t *testing.T) {
	t.Run("Test with slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		slice := []string{"test1", "test2", "test3"}
		result := newSliceElement(slice)
		if result == nil {
			t.Errorf("Expected a new slice, got nil")
		}
		newSlice, ok := result.(*[]string)
		if !ok {
			t.Errorf("Expected a *[]string, got %T", result)
		}
		if len(*newSlice) != 0 {
			t.Errorf("Expected a new slice with length 0, got %d", len(*newSlice))
		}
	})

	t.Run("Test with nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := newSliceElement(nil)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("Test with unsupported type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := newSliceElement(123)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}
