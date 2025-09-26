package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap_2(t *testing.T) {
	t.Run("TestUnwrap", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testErr := errors.New("test error")
		msg := &Error{
			Err: testErr,
		}
		unwrappedErr := msg.Unwrap()
		if unwrappedErr != testErr {
			t.Errorf("Expected %v, got %v", testErr, unwrappedErr)
		}
	})
}
