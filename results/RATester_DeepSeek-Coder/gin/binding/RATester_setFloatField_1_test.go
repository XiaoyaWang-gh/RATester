package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFloatField_1(t *testing.T) {
	type TestStruct struct {
		FloatField float64
	}

	t.Run("Test with valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.FloatField).Elem()
		err := setFloatField("123.456", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.FloatField != 123.456 {
			t.Errorf("Expected 123.456, got %v", testStruct.FloatField)
		}
	})

	t.Run("Test with empty input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.FloatField).Elem()
		err := setFloatField("", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.FloatField != 0.0 {
			t.Errorf("Expected 0.0, got %v", testStruct.FloatField)
		}
	})

	t.Run("Test with invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.FloatField).Elem()
		err := setFloatField("invalid", 64, field)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
