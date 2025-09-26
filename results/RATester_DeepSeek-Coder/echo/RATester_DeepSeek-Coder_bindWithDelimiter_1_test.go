package echo

import (
	"fmt"
	"testing"
)

func TestBindWithDelimiter_1(t *testing.T) {
	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"1,2,3"}
			}
			return nil
		},
		ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
			return fmt.Errorf("error: %v", message)
		},
	}

	t.Run("Test bindWithDelimiter with string slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []string
		b.bindWithDelimiter("ids", &dest, ",", true)
		if len(dest) != 3 {
			t.Errorf("Expected 3 values, got %d", len(dest))
		}
	})

	t.Run("Test bindWithDelimiter with bool slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []bool
		b.bindWithDelimiter("ids", &dest, ",", true)
		if len(dest) != 3 {
			t.Errorf("Expected 3 values, got %d", len(dest))
		}
	})

	t.Run("Test bindWithDelimiter with int slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int
		b.bindWithDelimiter("ids", &dest, ",", true)
		if len(dest) != 3 {
			t.Errorf("Expected 3 values, got %d", len(dest))
		}
	})

	t.Run("Test bindWithDelimiter with unsupported type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest struct{}
		b.bindWithDelimiter("ids", &dest, ",", true)
		if b.errors == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
