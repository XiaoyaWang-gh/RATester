package template

import (
	"fmt"
	htmltemplate "html/template"
	"testing"
)

func TestStringify_1(t *testing.T) {
	t.Run("Test with single string argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, _ := stringify("test")
		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})

	t.Run("Test with multiple arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, _ := stringify("test", 1, true)
		if result != "test 1 true" {
			t.Errorf("Expected 'test 1 true', got '%s'", result)
		}
	})

	t.Run("Test with nil argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, _ := stringify(nil)
		if result != "" {
			t.Errorf("Expected '', got '%s'", result)
		}
	})

	t.Run("Test with htmltemplate types", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, _ := stringify(htmltemplate.CSS("test"))
		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})

	t.Run("Test with multiple arguments and htmltemplate types", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, _ := stringify("test", htmltemplate.HTML("<test>"), 1, true)
		if result != "test <test> 1 true" {
			t.Errorf("Expected 'test <test> 1 true', got '%s'", result)
		}
	})
}
