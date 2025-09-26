package strings

import (
	"fmt"
	"testing"
)

func TestCount_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		count, err := ns.Count("hello", "hello world")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 1 {
			t.Errorf("Expected count to be 1, got %v", count)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Count(123, "hello world")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Count("hello", 123)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
