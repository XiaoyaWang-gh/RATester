package safe

import (
	"fmt"
	"testing"
)

func TestHTMLAttr_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		attr, err := ns.HTMLAttr("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if attr != "test" {
			t.Errorf("Expected 'test', got %v", attr)
		}
	})

	t.Run("int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.HTMLAttr(123)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.HTMLAttr(nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
