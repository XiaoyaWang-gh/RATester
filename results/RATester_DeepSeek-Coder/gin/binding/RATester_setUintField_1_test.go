package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetUintField_1(t *testing.T) {
	type TestStruct struct {
		Field uint64
	}

	t.Run("ValidInput", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setUintField("123", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.Field != 123 {
			t.Errorf("Expected Field to be 123, got %d", testStruct.Field)
		}
	})

	t.Run("EmptyInput", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setUintField("", 64, field)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if testStruct.Field != 0 {
			t.Errorf("Expected Field to be 0, got %d", testStruct.Field)
		}
	})

	t.Run("InvalidInput", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testStruct := TestStruct{}
		field := reflect.ValueOf(&testStruct.Field).Elem()
		err := setUintField("abc", 64, field)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if testStruct.Field != 0 {
			t.Errorf("Expected Field to be 0, got %d", testStruct.Field)
		}
	})
}
