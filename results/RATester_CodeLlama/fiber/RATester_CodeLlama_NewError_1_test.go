package fiber

import (
	"fmt"
	"testing"
)

func TestNewError_1(t *testing.T) {
	t.Parallel()
	t.Run("NewError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		code := 404
		message := "Not Found"
		err := NewError(code, message)
		if err.Code != code {
			t.Errorf("NewError() = %v, want %v", err.Code, code)
		}
		if err.Message != message {
			t.Errorf("NewError() = %v, want %v", err.Message, message)
		}
	})
}
