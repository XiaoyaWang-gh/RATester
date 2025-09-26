package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestCSS_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := template.CSS("test")
		actual, err := ns.CSS("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.CSS(make(chan int))
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
