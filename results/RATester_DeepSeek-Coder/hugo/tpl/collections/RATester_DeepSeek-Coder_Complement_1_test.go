package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComplement_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("Complement with slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Complement([]any{"a", "b", "c"}, "a", "b")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := []any{"c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Complement with array", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Complement([3]any{"a", "b", "c"}, "a", "b")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := [1]any{"c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Complement with less than two arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Complement("a")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Complement with arguments that are not slices or arrays", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Complement("a", "b")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
