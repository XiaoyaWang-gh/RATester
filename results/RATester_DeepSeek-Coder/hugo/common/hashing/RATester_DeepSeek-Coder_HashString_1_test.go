package hashing

import (
	"fmt"
	"testing"
)

func TestHashString_1(t *testing.T) {
	t.Run("Test with string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := HashString("test")
		expected := "3099084269960026951"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with multiple strings", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := HashString("test1", "test2", "test3")
		expected := "10141931000000000000"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with integers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := HashString(1, 2, 3)
		expected := "10141931000000000000"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with float", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := HashString(1.1, 2.2, 3.3)
		expected := "10141931000000000000"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
