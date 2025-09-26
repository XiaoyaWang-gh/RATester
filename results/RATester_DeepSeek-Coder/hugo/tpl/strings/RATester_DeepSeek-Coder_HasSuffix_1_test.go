package strings

import (
	"fmt"
	"testing"
)

func TestHasSuffix_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		suffix := "foo"
		s := "barfoo"
		expected := true

		result, err := ns.HasSuffix(s, suffix)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		suffix := 123
		s := "barfoo"
		expected := "cast.ToStringE: unable to cast 123 of type int to string"

		_, err := ns.HasSuffix(s, suffix)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != expected {
			t.Errorf("Expected error %q, got %q", expected, err.Error())
		}
	})
}
