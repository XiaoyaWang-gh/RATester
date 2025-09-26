package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetIntField_1(t *testing.T) {
	type TestStruct struct {
		Field int64
	}

	t.Run("Test with valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setIntField("123", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.Field != 123 {
			t.Errorf("Expected Field to be 123, got %v", testStruct.Field)
		}
	})

	t.Run("Test with invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setIntField("invalid", 64, field)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if testStruct.Field != 0 {
			t.Errorf("Expected Field to be 0, got %v", testStruct.Field)
		}
	})

	t.Run("Test with empty input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setIntField("", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.Field != 0 {
			t.Errorf("Expected Field to be 0, got %v", testStruct.Field)
		}
	})
}
