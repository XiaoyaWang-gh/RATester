package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseWhereArgs_1(t *testing.T) {
	t.Run("Test with one argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		arg := "test"
		mv, op, err := parseWhereArgs(arg)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if op != "" {
			t.Errorf("Expected empty operator, got %v", op)
		}
		if mv.Kind() != reflect.String {
			t.Errorf("Expected argument to be of type string, got %v", mv.Kind())
		}
	})

	t.Run("Test with two arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		arg1 := ">"
		arg2 := 10
		mv, op, err := parseWhereArgs(arg1, arg2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if op != ">" {
			t.Errorf("Expected operator to be '>', got %v", op)
		}
		if mv.Kind() != reflect.Int {
			t.Errorf("Expected argument to be of type int, got %v", mv.Kind())
		}
	})

	t.Run("Test with no arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, _, err := parseWhereArgs()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Test with more than two arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		arg1 := ">"
		arg2 := 10
		arg3 := "test"
		_, _, err := parseWhereArgs(arg1, arg2, arg3)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
