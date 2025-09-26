package collections

import (
	"fmt"
	"testing"
)

func TestEchoParam_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("TestEchoParam_Array", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		arr := []any{1, 2, 3, 4, 5}
		expected := 3
		result := ns.EchoParam(arr, 2)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestEchoParam_Slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		slice := []any{1, 2, 3, 4, 5}
		expected := 3
		result := ns.EchoParam(slice, 2)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestEchoParam_Map", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		m := map[string]any{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		}
		expected := 2
		result := ns.EchoParam(m, "key2")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestEchoParam_InvalidType", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		invalid := "invalid"
		expected := ""
		result := ns.EchoParam(invalid, 2)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestEchoParam_Nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var nilVar any
		expected := ""
		result := ns.EchoParam(nilVar, 2)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
