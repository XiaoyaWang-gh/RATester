package strings

import (
	"fmt"
	"testing"
)

func TestReplace_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Replace("hello world", "world", "Go")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello Go" {
			t.Errorf("Expected 'hello Go', got '%s'", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Replace(123, "world", "Go")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("limit", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Replace("hello world world", "world", "Go", 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello Go world" {
			t.Errorf("Expected 'hello Go world', got '%s'", result)
		}
	})
}
