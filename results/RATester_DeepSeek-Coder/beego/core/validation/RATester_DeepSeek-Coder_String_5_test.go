package validation

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	t.Run("TestString", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		e := &Error{
			Message: "Test Error",
		}
		expected := "Test Error"
		result := e.String()
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
