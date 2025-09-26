package web

import (
	"fmt"
	"testing"
)

func TestError_5(t *testing.T) {
	fd := &FlashData{Data: make(map[string]string)}

	t.Run("WithoutArgs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		msg := "error message"
		fd.Error(msg)
		if fd.Data["error"] != msg {
			t.Errorf("Expected '%s', got '%s'", msg, fd.Data["error"])
		}
	})

	t.Run("WithArgs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		msg := "error message %d"
		arg := 123
		expected := fmt.Sprintf(msg, arg)
		fd.Error(msg, arg)
		if fd.Data["error"] != expected {
			t.Errorf("Expected '%s', got '%s'", expected, fd.Data["error"])
		}
	})
}
