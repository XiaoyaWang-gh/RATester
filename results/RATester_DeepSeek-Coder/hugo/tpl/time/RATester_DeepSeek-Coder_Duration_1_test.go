package time

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid unit and number", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		duration, err := ns.Duration("seconds", 5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if duration != 5*time.Second {
			t.Errorf("Expected duration to be 5 seconds, got %v", duration)
		}
	})

	t.Run("invalid unit", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Duration("invalid", 5)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("invalid number", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Duration("seconds", "five")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
