package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetError_1(t *testing.T) {
	t.Run("Test setError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{}
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")

		b.setError(err1)
		if len(b.errors) != 1 || b.errors[0] != err1 {
			t.Errorf("Expected error %v, got %v", err1, b.errors[0])
		}

		b.setError(err2)
		if len(b.errors) != 2 || b.errors[1] != err2 {
			t.Errorf("Expected error %v, got %v", err2, b.errors[1])
		}
	})
}
