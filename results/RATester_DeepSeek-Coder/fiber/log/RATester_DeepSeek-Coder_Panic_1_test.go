package log

import (
	"fmt"
	"testing"
)

func TestPanic_1(t *testing.T) {
	t.Run("should panic with a string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			} else if r != "test panic" {
				t.Errorf("Expected panic message to be 'test panic', got '%v'", r)
			}
		}()
		Panic("test panic")
	})

	t.Run("should panic with an integer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			} else if r != 123 {
				t.Errorf("Expected panic message to be 123, got '%v'", r)
			}
		}()
		Panic(123)
	})
}
