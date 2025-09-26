package css

import (
	"fmt"
	"testing"
)

func TestQuoted_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs := ns.Quoted("test")
		if string(qs) != "\"test\"" {
			t.Errorf("Expected \"test\", got %s", qs)
		}
	})

	t.Run("int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs := ns.Quoted(123)
		if string(qs) != "\"123\"" {
			t.Errorf("Expected \"123\", got %s", qs)
		}
	})

	t.Run("float", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs := ns.Quoted(123.456)
		if string(qs) != "\"123.456\"" {
			t.Errorf("Expected \"123.456\", got %s", qs)
		}
	})

	t.Run("bool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs := ns.Quoted(true)
		if string(qs) != "\"true\"" {
			t.Errorf("Expected \"true\", got %s", qs)
		}
	})
}
