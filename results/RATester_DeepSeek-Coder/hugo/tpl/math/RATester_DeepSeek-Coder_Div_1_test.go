package math

import (
	"fmt"
	"testing"
)

func TestDiv_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("TestDiv_Success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Div(10, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 5 {
			t.Errorf("Expected 5, got %v", result)
		}
	})

	t.Run("TestDiv_Error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Div("10", 2)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
