package metric

import (
	"fmt"
	"testing"
)

func TestNewDateCounter_1(t *testing.T) {
	t.Run("Test with positive reserveDays", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		dc := NewDateCounter(7)
		if dc == nil {
			t.Errorf("Expected a DateCounter, got nil")
		}
	})

	t.Run("Test with zero reserveDays", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		dc := NewDateCounter(0)
		if dc == nil {
			t.Errorf("Expected a DateCounter, got nil")
		}
	})

	t.Run("Test with negative reserveDays", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		dc := NewDateCounter(-7)
		if dc == nil {
			t.Errorf("Expected a DateCounter, got nil")
		}
	})
}
