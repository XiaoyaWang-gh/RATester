package template

import (
	"fmt"
	"testing"
)

func TestEvalArgs_1(t *testing.T) {
	t.Run("SingleStringArg", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := evalArgs("test")
		expected := "test"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("MultipleArgs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := evalArgs("test", 1, true)
		expected := "test 1 true"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("SingleNonStringArg", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := evalArgs(1)
		expected := "1"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
