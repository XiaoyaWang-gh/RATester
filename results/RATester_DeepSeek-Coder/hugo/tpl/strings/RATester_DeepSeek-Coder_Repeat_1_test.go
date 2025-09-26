package strings

import (
	"fmt"
	"testing"
)

func TestRepeat_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.Repeat(3, "abc")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if got != "abcabcabc" {
			t.Errorf("Expected 'abcabcabc', got '%s'", got)
		}
	})

	t.Run("negative count", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Repeat(-1, "abc")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "strings: negative Repeat count" {
			t.Errorf("Expected 'strings: negative Repeat count', got '%s'", err.Error())
		}
	})

	t.Run("non-string second argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Repeat(3, 123)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "cast: unable to cast 123 of type int to string" {
			t.Errorf("Expected 'cast: unable to cast 123 of type int to string', got '%s'", err.Error())
		}
	})

	t.Run("non-int first argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Repeat("abc", "123")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "cast: unable to cast 123 of type string to int" {
			t.Errorf("Expected 'cast: unable to cast 123 of type string to int', got '%s'", err.Error())
		}
	})
}
