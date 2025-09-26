package css

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/css"
)

func TestUnquoted_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted("test")
		expected := css.UnquotedString("test")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(123)
		expected := css.UnquotedString("123")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("float", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(123.456)
		expected := css.UnquotedString("123.456")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("bool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(true)
		expected := css.UnquotedString("true")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
