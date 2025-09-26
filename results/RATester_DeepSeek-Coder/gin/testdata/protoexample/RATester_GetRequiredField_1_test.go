package protoexample

import (
	"fmt"
	"testing"
)

func TestGetRequiredField_1(t *testing.T) {
	t.Run("Test with nil Test_OptionalGroup", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var x *Test_OptionalGroup
		result := x.GetRequiredField()
		expected := ""
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with non-nil Test_OptionalGroup with nil RequiredField", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		x := &Test_OptionalGroup{}
		result := x.GetRequiredField()
		expected := ""
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with non-nil Test_OptionalGroup with non-nil RequiredField", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		value := "test"
		x := &Test_OptionalGroup{RequiredField: &value}
		result := x.GetRequiredField()
		expected := value
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})
}
