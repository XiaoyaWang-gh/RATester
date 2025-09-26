package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBoolField_1(t *testing.T) {
	type TestStruct struct {
		BoolField bool
	}

	t.Run("Test with valid string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.BoolField).Elem()
		err := setBoolField("true", field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !testStruct.BoolField {
			t.Errorf("Expected BoolField to be true, got false")
		}
	})

	t.Run("Test with invalid string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.BoolField).Elem()
		err := setBoolField("invalid", field)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if testStruct.BoolField {
			t.Errorf("Expected BoolField to be false, got true")
		}
	})

	t.Run("Test with empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.BoolField).Elem()
		err := setBoolField("", field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !testStruct.BoolField {
			t.Errorf("Expected BoolField to be true, got false")
		}
	})
}
