package validation

import (
	"fmt"
	"testing"
)

func TestMessage_1(t *testing.T) {
	t.Run("should set error message with fmt.Sprintf", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r := &Result{
			Error: &Error{},
			Ok:    false,
		}
		message := "test message %s"
		arg := "arg"
		expected := fmt.Sprintf(message, arg)
		r.Message(message, arg)
		if r.Error.Message != expected {
			t.Errorf("expected %s, got %s", expected, r.Error.Message)
		}
	})

	t.Run("should set error message without fmt.Sprintf", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r := &Result{
			Error: &Error{},
			Ok:    false,
		}
		message := "test message"
		r.Message(message)
		if r.Error.Message != message {
			t.Errorf("expected %s, got %s", message, r.Error.Message)
		}
	})
}
